from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By
from selenium.webdriver.support.expected_conditions import presence_of_element_located
from selenium.common.exceptions import TimeoutException, NoSuchElementException
from selenium.webdriver.support import expected_conditions
import time
import kanjiData

WANIKANI = 'https://www.wanikani.com/login'

class WanikaniParser:
    def __init__(self, headless: bool = False):
        opts = webdriver.firefox.options.Options()
        opts.headless = headless
        self.__driver = webdriver.Firefox(options=opts)
        self.__waitOneSecond = WebDriverWait(self.__driver, 1)
        self.__wait = WebDriverWait(self.__driver, 10)

    def __enter__(self):
        self.__driver.get(WANIKANI)
        return self
    
    def __exit__(self, exc_type, exc_value, exc_traceback):
        self.__driver.close()

    def Login(self, email: str, password: str):
        self.__driver.find_element(By.ID, 'user_login').send_keys(email) #user_login
        self.__driver.find_element(By.ID, 'user_password').send_keys(password + Keys.ENTER) #user_password
        try:
            error =  self.__waitOneSecond.until(lambda d: d.find_element(By.CLASS_NAME, 'alert-error'))
        except TimeoutException:
            return
        if error:
            raise ValueError("Invalid login or password")

    def GetLevelsButtons(self) -> list:
        levels = self.__wait.until(lambda d: d.find_elements(By.CLASS_NAME, 'sitemap__pages--levels'))
        levelsList = []
        for level in levels:
            filter = lambda d: level.find_elements(By.TAG_NAME, 'li')
            levelsList += self.__GetLinks(filter)
        return levelsList

    def GetAllKanjiFromLevel(self, level, kanjiFilter) -> kanjiData.Kanji:
        self.__GetToLevel(level)
        filter = lambda d: d.find_elements(By.XPATH, "//*[contains(@class, 'kanji-') and contains(@class, 'character-item')]")
        kanjiList = self.__GetLinks(filter)
        for kanji in kanjiList:
            if kanji in kanjiFilter: 
                continue
            self.__driver.get(kanji)
            yield self.__GetKanjiData(), kanji

    def GetAllWordsFromLevel(self, level, wordFilter) -> kanjiData.Word:
        self.__GetToLevel(level)
        filter = lambda d: d.find_elements(By.XPATH, "//*[contains(@class, 'vocabulary-') and contains(@class, 'character-item')]")
        wordsList = self.__GetLinks(filter)
        for word in wordsList:
            if word in wordFilter: 
                continue
            self.__driver.get(word)
            yield self.__GetWordData(), word

    def __GetKanjiData(self) -> kanjiData.Kanji:
        kanji = kanjiData.Kanji()
        try:
            kanji.name = self.__wait.until(lambda d: d.find_element(By.CLASS_NAME, 'kanji-icon')).text
        except TimeoutException:
            return None
        kanji.level = int(self.__driver.find_element(By.CLASS_NAME, 'level-icon').text)

        meanings = self.__driver.find_element(By.ID, 'meaning').find_elements(By.CLASS_NAME, 'alternative-meaning')
        kanji.primary = meanings[0].find_element(By.TAG_NAME, 'p').text
        if (len(meanings) > 2):
            kanji.alternatives = meanings[1].find_element(By.TAG_NAME, 'p').text.strip().split(', ')

        readings = self.__driver.find_element(By.ID, 'reading').find_elements(By.CLASS_NAME, 'span4')
        kanji.onyomi = readings[0].find_element(By.TAG_NAME, 'p').text.strip().split(', ')
        kanji.kunyomi = readings[1].find_element(By.TAG_NAME, 'p').text.strip().split(', ')

        try:
            kanji.progress = self.__driver.find_element(By.CLASS_NAME, 'srs').find_element(By.CSS_SELECTOR, "div[style*='display: inline-block;']").text.strip() 
        except NoSuchElementException:
            pass
        return kanji

    def __GetWordData(self) -> kanjiData.Word:
        word = kanjiData.Word()
        try:
            word.name = self.__wait.until(lambda d: d.find_element(By.CLASS_NAME, 'vocabulary-icon')).text
        except TimeoutException:
            return None
        word.level = int(self.__driver.find_element(By.CLASS_NAME, 'level-icon').text)

        meanings = self.__driver.find_element(By.ID, 'meaning').find_elements(By.CLASS_NAME, 'alternative-meaning')
        word.primary = meanings[0].find_element(By.TAG_NAME, 'p').text
        alt = meanings[1].find_element(By.TAG_NAME, 'h2').text.lower()
        if "alter" in alt:
            word.alternatives = meanings[1].find_element(By.TAG_NAME, 'p').text.strip().split(', ')
        word.wordType = meanings[-1].find_element(By.TAG_NAME, 'p').text.strip().split(', ')

        readings = self.__driver.find_element(By.ID, 'reading').find_elements(By.CLASS_NAME, 'pronunciation-variant')
        word.reading = [reading.text for reading in readings]

        context = self.__driver.find_element(By.ID, 'context')
        # wordTypes = context.find_elements(By.CLASS_NAME, 'word-type')
        # word.wordType = [wordType.text for wordType in wordTypes]

        sentencesGroup = context.find_elements(By.CLASS_NAME, 'context-sentence-group')
        sentencesText = [sentences.find_elements(By.TAG_NAME, 'p') for sentences in sentencesGroup]
        sentences = []
        for s in sentencesText:
            sentences.append(kanjiData.Sentence(s[0].text, s[1].text))
        word.sentences = sentences

        composition = self.__driver.find_element(By.ID, 'components').find_elements(By.CLASS_NAME, 'character')
        word.composition = [comp.text for comp in composition]
        try:
            word.progress = self.__driver.find_element(By.CLASS_NAME, 'srs').find_element(By.CSS_SELECTOR, "div[style*='display: inline-block;']").text.strip() 
        except NoSuchElementException:
            pass

        return word


    def __GetLinks(self, filter) -> list:
        items = self.__wait.until(filter)
        items = [item.find_element(By.TAG_NAME, 'a').get_attribute('href') for item in items]
        return items

    def __GetToLevel(self, level):
        self.__driver.get(level)  