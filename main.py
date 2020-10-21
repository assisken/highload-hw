import os

from app.create_app import create_app

DEBUG = os.getenv("DEBUG", "").lower() in ("true", "1")
app = create_app()

if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=DEBUG, port=80)
