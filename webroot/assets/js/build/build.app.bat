node r.js -o build.app.js

java -jar closure-compiler-v20180204.jar --js=../dev/compressor.js --js_output_file=../pro/compressor.js

java -jar closure-compiler-v20180204.jar --js=../pro/page/home/Default.js --js_output_file=../pro/page/home/Default.min.js
del /q ..\pro\page\home\Default.js

java -jar closure-compiler-v20180204.jar --js=../pro/page/home/Cell.js --js_output_file=../pro/page/home/Cell.min.js
del /q ..\pro\page\home\Cell.js

java -jar closure-compiler-v20180204.jar --js=../pro/page/home/Cells.js --js_output_file=../pro/page/home/Cells.min.js
del /q ..\pro\page\home\Cells.js

java -jar closure-compiler-v20180204.jar --js=../pro/page/home/Login.js --js_output_file=../pro/page/home/Login.min.js
del /q ..\pro\page\home\Login.js

del /q ..\pro\mylib
del /q ..\pro\build.txt
