from contextlib import contextmanager

from db.base import Session

class AbstractSession:
    @contextmanager
    def session_scope(self):
        session = Session()
        try:
            yield session
            session.commit()
        except:
            session.rollback()
            raise
        finally:
            session.close()
