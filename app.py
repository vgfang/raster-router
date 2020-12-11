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
	fid = request.cookies.get('fid')
	if "/" in fid :
		return "Unsafe modified file ID. Could cause overwrites in other folders."

	# generate present and next version
	suffix1 = "_pres.png"
	suffix2 = "_next.png"
	fpath1 = (tempDir + fid + suffix1).encode()
	fpath2 = (tempDir + fid + suffix2).encode()
	fpath1_g = GoString(fpath1, len(fpath1)) # generating two instead of copying
	fpath2_g = GoString(fpath2, len(fpath2))

	imgproc.generate_blank.argtypes = [GoString, ct.c_longlong, ct.c_longlong]
	imgproc.generate_blank(fpath1_g, width, height)
	imgproc.generate_blank(fpath2_g, width, height)
	return "Successful Blank Generation."


@app.route('/create/move_on', methods = ['POST'])
def move_on():
	fid = request.cookies.get('fid')
	if "/" in fid :
		return "Unsafe modified file ID. Could cause overwrites in other folders."
	fpre = tempDir+fid
	# os.rename(fpre+"_pres.png", fpre+"_prev.png") # rename present to previous
	shutil.copyfile(fpre+"_next.png", fpre+"_pres.png") # copy next to present
	return "Successful Move On."

@app.route('/create/run_command', methods=['POST'])
def run_command():
	fid = request.cookies.get('fid')
	if "/" in fid :
		return "Unsafe modified file ID. Could cause overwrites in other folders."
	rs_comm = request.form.get("rs_comm")
	if len(rs_comm) < 6:
		return "Invalid command String."

	# generate next version
	suffix1 = "_pres.png"
	suffix2 = "_next.png"
	fpath1 = (tempDir + fid + suffix1).encode()
	fpath2 = (tempDir + fid + suffix2).encode()
	fpath1_g = GoString(fpath1, len(fpath1)) # generating two instead of copying
	fpath2_g = GoString(fpath2, len(fpath2))
	rs_commGstr = GoString(rs_comm.encode(), len(rs_comm.encode()))

	imgproc.run_command_str.argtypes = [GoString, GoString, GoString]
	imgproc.run_command_str(fpath1_g, fpath2_g, rs_commGstr)

	return "Successful Draw Command Next Generation."

@app.route('/create/new_argimg', methods=['POST'])
def new_argimg():
	return "none"

if __name__ == "__main__":
	app.run(debug=True, host= '0.0.0.0')
