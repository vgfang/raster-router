from flask import Flask, render_template, url_for, request, make_response, after_this_request, redirect
from ctypes import * # foreign function library
import uuid # for filename unique id

app = Flask(__name__)
app.config['SQLAlLCHEMY_DATABASE_URL'] = "sqlite:///test.db"

# set unique cookie id for filename before every page
@app.before_request
def check_fid():
	if not request.cookies.get('fid'):
		@after_this_request
		def set_fid(response):
			new_id = uuid.uuid4().hex
			response.set_cookie('fid', new_id)
			return response

@app.route('/')
def show_index():
	return render_template("index.html")

@app.route('/create')
def show_create():
	if not request.cookies.get('fid'):
		return redirect("/")
	return render_template("create.html",fid=request.cookies.get('fid'))


if __name__ == "__main__":
	app.run(debug=True, host= '0.0.0.0')
