import os

from app.create_app import create_app

DEBUG = bool(os.getenv("DEBUG", ""))
app = create_app()
app.run(host="0.0.0.0", debug=DEBUG, port=8000)
