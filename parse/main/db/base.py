from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

import os
from pathlib import Path
import configparser

configName = 'config.ini'
configPath = os.path.join(Path(__file__).parents[1], configName)

config = configparser.ConfigParser()
config.read(configPath)

dbPath = config.get('DATABASE' ,'Path')

Engine = create_engine(dbPath)
Session = sessionmaker(bind=Engine)

Base = declarative_base()