from pytest import fixture

from app.app import create_app


@fixture
def app():
    app = create_app()
    return app
