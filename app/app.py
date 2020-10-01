import os

from flask import Flask

from app.api.v1.forecast import api

DEBUG = bool(os.getenv("DEBUG", ""))


def create_app():
    app = Flask(__name__)
    app.register_blueprint(api, url_prefix="/v1")
    return app


if __name__ == "__main__":
    app = create_app()
    app.run(host="0.0.0.0", debug=DEBUG, port=8000)
