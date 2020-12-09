import ctypes as ct # foreign function library

imgproc = ct.cdll.LoadLibrary('./imgproc.so') # golang functions via shared libary
class GoString(ct.Structure):
    _fields_ = [("p", ct.c_char_p), ("n", ct.c_longlong)]


imgproc.generate_blank.argtypes = [GoString, ct.c_longlong, ct.c_longlong]
imgproc.generate_blank(GoString("apple.png".encode(),len(b"apple.png")),300,300)
