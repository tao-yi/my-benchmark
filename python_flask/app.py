from flask import Flask
import logging

log = logging.getLogger('werkzeug')
log.disabled = True
app = Flask(__name__)


@app.route("/", methods=["GET"])
def hello():
    return {"message": "hello world"}


if __name__ == '__main__':
    app.run(host="localhost", port=3999, debug=False)
