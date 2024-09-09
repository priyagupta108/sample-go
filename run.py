import ctypes
import time
import os

# Load the shared library from an environment variable
lib_path = os.getenv('RUNPY_LIB_PATH', './lib.so')
lib = ctypes.cdll.LoadLibrary(lib_path)

# Start timing
start_time = time.time()

# Call cNoop 1 million times
for _ in range(1000000):
    lib.cNoop()

# End timing
end_time = time.time()

print(f"Made 1000000 calls in {end_time - start_time:.2f}s")