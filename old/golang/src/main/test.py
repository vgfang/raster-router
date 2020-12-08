from ctypes import *

def test_function():
	'''
	aproc = cdll.LoadLibrary("./main..so")
	aproc.Add.argtypes = [c_longlong, c_longlong]
	print(aproc.Add(12,99))

	filename = b"test"
	aproc.Load.argtypes = [c_char_p]
	aproc.Load.restype = c_char_p
	print(aproc.Load(filename).decode("utf-8"))
	'''

test_function()
