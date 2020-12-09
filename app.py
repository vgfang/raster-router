from flask import Flask, render_template, url_for, request, make_response, after_this_request, redirect
import ctypes as ct # foreign function library
import uuid # for filename unique id
import os # for file deleting
import shutil  # for file copying

app = Flask(__name__)
app.config['SQLAlLCHEMY_DATABASE_URL'] = "sqlite:///test.db"

imgproc = ct.cdll.LoadLibrary('./imgproc.so') # golang functions via shared libary
tempDir = "./static/img/user/"
class GoString(ct.Structure): # structure for calling golang string
    _fields_ = [("p", ct.c_char_p), ("n", ct.c_longlong)]

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

# calls golang to generate 2 blank images
@app.route('/create/generate_blank', methods = ['POST'])
def generate_blank():
	height = int(request.form.get("height"))
	width = int(request.form.get("width"))

	# generate present and past version
	suffix1 = "_pres.png"
	suffix2 = "_next.png"
	fnstr1 = (tempDir + request.cookies.get('fid') + suffix1).encode()
	fnstr2 = (tempDir + request.cookies.get('fid') + suffix2).encode()
	fnGstr1 = GoString(fnstr1, len(fnstr1)) # generating two instead of copying
	fnGstr2 = GoString(fnstr2, len(fnstr2))

	imgproc.generate_blank.argtypes = [GoString, ct.c_longlong, ct.c_longlong]
	imgproc.generate_blank(fnGstr1, width, height)
	imgproc.generate_blank(fnGstr2, width, height)
	return "Successful Blank Generation."


@app.route('/create/move_on', methods = ['POST'])
def move_on():
	fid = request.cookies.get('fid')
	fpre = tempDir+fid
	# os.rename(fpre+"_pres.png", fpre+"_prev.png") # rename present to previous
	shutil.copyfile(fpre+"_next", fpre+"_prev.png") # copy next to present



if __name__ == "__main__":
	app.run(debug=True, host= '0.0.0.0')
