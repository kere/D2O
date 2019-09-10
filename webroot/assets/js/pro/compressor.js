/*
 Compressor.js v1.0.5
 https://fengyuanchen.github.io/compressorjs

 Copyright 2018-present Chen Fengyuan
 Released under the MIT license

 Date: 2019-01-23T10:53:08.724Z
*/
var $jscomp=$jscomp||{};$jscomp.scope={};$jscomp.owns=function(a,f){return Object.prototype.hasOwnProperty.call(a,f)};$jscomp.assign="function"==typeof Object.assign?Object.assign:function(a,f){for(var e=1;e<arguments.length;e++){var h=arguments[e];if(h)for(var l in h)$jscomp.owns(h,l)&&(a[l]=h[l])}return a};$jscomp.ASSUME_ES5=!1;$jscomp.ASSUME_NO_NATIVE_MAP=!1;$jscomp.ASSUME_NO_NATIVE_SET=!1;
$jscomp.defineProperty=$jscomp.ASSUME_ES5||"function"==typeof Object.defineProperties?Object.defineProperty:function(a,f,e){a!=Array.prototype&&a!=Object.prototype&&(a[f]=e.value)};$jscomp.getGlobal=function(a){return"undefined"!=typeof window&&window===a?a:"undefined"!=typeof global&&null!=global?global:a};$jscomp.global=$jscomp.getGlobal(this);
$jscomp.polyfill=function(a,f,e,h){if(f){e=$jscomp.global;a=a.split(".");for(h=0;h<a.length-1;h++){var l=a[h];l in e||(e[l]={});e=e[l]}a=a[a.length-1];h=e[a];f=f(h);f!=h&&null!=f&&$jscomp.defineProperty(e,a,{configurable:!0,writable:!0,value:f})}};$jscomp.polyfill("Object.assign",function(a){return a||$jscomp.assign},"es6","es3");$jscomp.SYMBOL_PREFIX="jscomp_symbol_";$jscomp.initSymbol=function(){$jscomp.initSymbol=function(){};$jscomp.global.Symbol||($jscomp.global.Symbol=$jscomp.Symbol)};
$jscomp.Symbol=function(){var a=0;return function(f){return $jscomp.SYMBOL_PREFIX+(f||"")+a++}}();$jscomp.initSymbolIterator=function(){$jscomp.initSymbol();var a=$jscomp.global.Symbol.iterator;a||(a=$jscomp.global.Symbol.iterator=$jscomp.global.Symbol("iterator"));"function"!=typeof Array.prototype[a]&&$jscomp.defineProperty(Array.prototype,a,{configurable:!0,writable:!0,value:function(){return $jscomp.arrayIterator(this)}});$jscomp.initSymbolIterator=function(){}};
$jscomp.arrayIterator=function(a){var f=0;return $jscomp.iteratorPrototype(function(){return f<a.length?{done:!1,value:a[f++]}:{done:!0}})};$jscomp.iteratorPrototype=function(a){$jscomp.initSymbolIterator();a={next:a};a[$jscomp.global.Symbol.iterator]=function(){return this};return a};
$jscomp.iteratorFromArray=function(a,f){$jscomp.initSymbolIterator();a instanceof String&&(a+="");var e=0,h={next:function(){if(e<a.length){var l=e++;return{value:f(l,a[l]),done:!1}}h.next=function(){return{done:!0,value:void 0}};return h.next()}};h[Symbol.iterator]=function(){return h};return h};$jscomp.polyfill("Array.prototype.keys",function(a){return a?a:function(){return $jscomp.iteratorFromArray(this,function(a){return a})}},"es6","es3");
(function(a){"object"===typeof exports&&"undefined"!==typeof module?module.exports=a():"function"===typeof define&&define.amd?define([],a):("undefined"!==typeof window?window:"undefined"!==typeof global?global:"undefined"!==typeof self?self:this).Compressor=a()})(function(){return function(){function a(f,e,h){function l(y,v){if(!e[y]){if(!f[y]){var z="function"==typeof require&&require;if(!v&&z)return z(y,!0);if(I)return I(y,!0);v=Error("Cannot find module '"+y+"'");throw v.code="MODULE_NOT_FOUND",
v;}v=e[y]={exports:{}};f[y][0].call(v.exports,function(a){return l(f[y][1][a]||a)},v,v.exports,a,f,e,h)}return e[y].exports}for(var I="function"==typeof require&&require,z=0;z<h.length;z++)l(h[z]);return l}return a}()({1:[function(a,f,e){(function(a,l){"object"===typeof e&&"undefined"!==typeof f?f.exports=l():(a=a||self,a.Compressor=l())})(this,function(){function a(a,b){for(var g=0;g<b.length;g++){var c=b[g];c.enumerable=c.enumerable||!1;c.configurable=!0;"value"in c&&(c.writable=!0);Object.defineProperty(a,
c.key,c)}}function f(F,b,g){b&&a(F.prototype,b);g&&a(F,g);return F}function e(){e=Object.assign||function(a){for(var b=1;b<arguments.length;b++){var g=arguments[b],c;for(c in g)Object.prototype.hasOwnProperty.call(g,c)&&(a[c]=g[c])}return a};return e.apply(this,arguments)}function z(a){for(var b=1;b<arguments.length;b++){var g=null!=arguments[b]?arguments[b]:{},c=Object.keys(g);"function"===typeof Object.getOwnPropertySymbols&&(c=c.concat(Object.getOwnPropertySymbols(g).filter(function(b){return Object.getOwnPropertyDescriptor(g,
b).enumerable})));c.forEach(function(b){var c=g[b];b in a?Object.defineProperty(a,b,{value:c,enumerable:!0,configurable:!0,writable:!0}):a[b]=c})}return a}function y(a){return Array.from?Array.from(a):N.call(a)}function v(a){a=G.test(a)?a.substr(6):"";"jpeg"===a&&(a="jpg");return".".concat(a)}function H(a){var b=1<arguments.length&&void 0!==arguments[1]?arguments[1]:1E11;return O.test(a)?Math.round(a*b)/b:a}var M=function(a,b){return b={exports:{}},a(b,b.exports),b.exports}(function(a){"undefined"!==
typeof window&&function(b){var g=b.HTMLCanvasElement&&b.HTMLCanvasElement.prototype,c;if(c=b.Blob)try{c=!!new Blob}catch(P){c=!1}var d=c;if(c=d&&b.Uint8Array)try{c=100===(new Blob([new Uint8Array(100)])).size}catch(P){c=!1}var f=c,e=b.BlobBuilder||b.WebKitBlobBuilder||b.MozBlobBuilder||b.MSBlobBuilder,h=/^data:((.*?)(;charset=.*?)?)(;base64)?,/,B=(d||e)&&b.atob&&b.ArrayBuffer&&b.Uint8Array&&function(b){var a;var g=b.match(h);if(!g)throw Error("invalid data URI");var c=g[2]?g[1]:"text/plain"+(g[3]||
";charset=US-ASCII");var m=!!g[4];b=b.slice(g[0].length);b=m?atob(b):decodeURIComponent(b);m=new ArrayBuffer(b.length);g=new Uint8Array(m);for(a=0;a<b.length;a+=1)g[a]=b.charCodeAt(a);if(d)return new Blob([f?g:m],{type:c});b=new e;b.append(m);return b.getBlob(c)};b.HTMLCanvasElement&&!g.toBlob&&(g.mozGetAsFile?g.toBlob=function(b,a,c){var d=this;setTimeout(function(){c&&g.toDataURL&&B?b(B(d.toDataURL(a,c))):b(d.mozGetAsFile("blob",a))})}:g.toDataURL&&B&&(g.toBlob=function(b,a,g){var c=this;setTimeout(function(){b(B(c.toDataURL(a,
g)))})}));a.exports?a.exports=B:b.dataURLtoBlob=B}(window)}),J={strict:!0,checkOrientation:!0,maxWidth:Infinity,maxHeight:Infinity,minWidth:0,minHeight:0,width:void 0,height:void 0,quality:.8,mimeType:"auto",convertSize:5E6,beforeDraw:null,drew:null,success:null,error:null},C="undefined"!==typeof window?window:{},N=Array.prototype.slice,G=/^image\/.+$/,K=String.fromCharCode,Q=C.btoa,O=/\.\d*(?:0|9){12}\d*$/,R=C.ArrayBuffer,L=C.FileReader,A=C.URL||C.webkitURL,S=/\.\w+$/,T=C.Compressor;return function(){function a(b,
g){if(!(this instanceof a))throw new TypeError("Cannot call a class as a function");this.file=b;this.image=new Image;this.options=z({},J,g);this.aborted=!1;this.result=null;this.init()}f(a,[{key:"init",value:function(){var b=this,a=this.file,c=this.options;var d="undefined"===typeof Blob?!1:a instanceof Blob||"[object Blob]"===Object.prototype.toString.call(a);if(d){var f=a.type;if(G.test(f))if(A&&L)if(R||(c.checkOrientation=!1),A&&!c.checkOrientation)this.load({url:A.createObjectURL(a)});else{d=
new L;var h=c.checkOrientation&&"image/jpeg"===f;this.reader=d;d.onload=function(g){var c=g.target.result;g={};if(h){var d=new DataView(c);try{var k,q;if(255===d.getUint8(0)&&216===d.getUint8(1))for(var n=d.byteLength,m=2;m+1<n;){if(255===d.getUint8(m)&&225===d.getUint8(m+1)){var t=m;break}m+=1}if(t){n=t+10;t+=4;m=4;var u="",p;m+=t;for(p=t;p<m;p+=1)u+=K(d.getUint8(p));if("Exif"===u){var l=d.getUint16(n);if(((k=18761===l)||19789===l)&&42===d.getUint16(n+2,k)){var w=d.getUint32(n+4,k);8<=w&&(q=n+w)}}}if(q){var r=
d.getUint16(q,k),E;for(E=0;E<r;E+=1){var v=q+12*E+2;if(274===d.getUint16(v,k)){v+=8;var D=d.getUint16(v,k);d.setUint16(v,1,k);break}}}}catch(U){D=1}if(1<D||!A){k=[];for(c=new Uint8Array(c);0<c.length;)k.push(K.apply(null,y(c.subarray(0,8192)))),c=c.subarray(8192);c="data:".concat(f,";base64,").concat(Q(k.join("")));g.url=c;if(1<D){c=e;k=0;p=q=1;switch(D){case 2:q=-1;break;case 3:k=-180;break;case 4:p=-1;break;case 5:k=90;p=-1;break;case 6:k=90;break;case 7:k=90;q=-1;break;case 8:k=-90}c(g,{rotate:k,
scaleX:q,scaleY:p})}}else g.url=A.createObjectURL(a)}else g.url=c;b.load(g)};d.onabort=function(){b.fail(Error("Aborted to read the image with FileReader."))};d.onerror=function(){b.fail(Error("Failed to read the image with FileReader."))};d.onloadend=function(){b.reader=null};h?d.readAsArrayBuffer(a):d.readAsDataURL(a)}else this.fail(Error("The current browser does not support image compression."));else this.fail(Error("The first argument must be an image File or Blob object."))}else this.fail(Error("The first argument must be a File or Blob object."))}},
{key:"load",value:function(b){var a=this,c=this.file,d=this.image;d.onload=function(){a.draw(z({},b,{naturalWidth:d.naturalWidth,naturalHeight:d.naturalHeight}))};d.onabort=function(){a.fail(Error("Aborted to load the image."))};d.onerror=function(){a.fail(Error("Failed to load the image."))};d.alt=c.name;d.src=b.url}},{key:"draw",value:function(a){var b=this,c=a.naturalWidth,d=a.naturalHeight,f=a.rotate;f=void 0===f?0:f;var e=a.scaleX;e=void 0===e?1:e;a=a.scaleY;var h=void 0===a?1:a,l=this.file,
v=this.image;a=this.options;var k=document.createElement("canvas"),q=k.getContext("2d"),n=c/d,m=90===Math.abs(f)%180,t=Math.max(a.maxWidth,0)||Infinity,u=Math.max(a.maxHeight,0)||Infinity,p=Math.max(a.minWidth,0)||0,x=Math.max(a.minHeight,0)||0,w=Math.max(a.width,0)||c,r=Math.max(a.height,0)||d;m&&(u=[u,t],t=u[0],u=u[1],x=[x,p],p=x[0],x=x[1],r=[r,w],w=r[0],r=r[1]);Infinity>t&&Infinity>u?u*n>t?u=t/n:t=u*n:Infinity>t?u=t/n:Infinity>u&&(t=u*n);0<p&&0<x?x*n>p?x=p/n:p=x*n:0<p?x=p/n:0<x&&(p=x*n);r*n>w?
r=w/n:w=r*n;w=Math.floor(H(Math.min(Math.max(w,p),t)));r=Math.floor(H(Math.min(Math.max(r,x),u)));n=-w/2;t=-r/2;p=w;u=r;m&&(m=[r,w],w=m[0],r=m[1]);k.width=w;k.height=r;G.test(a.mimeType)||(a.mimeType=l.type);m="transparent";l.size>a.convertSize&&"image/png"===a.mimeType&&(m="#fff",a.mimeType="image/jpeg");q.fillStyle=m;q.fillRect(0,0,w,r);a.beforeDraw&&a.beforeDraw.call(this,q,k);this.aborted||(q.save(),q.translate(w/2,r/2),q.rotate(f*Math.PI/180),q.scale(e,h),q.drawImage(v,n,t,p,u),q.restore(),a.drew&&
a.drew.call(this,q,k),this.aborted||(f=function(a){b.aborted||b.done({naturalWidth:c,naturalHeight:d,result:a})},k.toBlob?k.toBlob(f,a.mimeType,a.quality):f(M(k.toDataURL(a.mimeType,a.quality)))))}},{key:"done",value:function(a){var b=a.naturalWidth,c=a.naturalHeight;a=a.result;var d=this.file,f=this.image,e=this.options;A&&!e.checkOrientation&&A.revokeObjectURL(f.src);a?e.strict&&a.size>d.size&&e.mimeType===d.type&&!(e.width>b||e.height>c||e.minWidth>b||e.minHeight>c)?a=d:(b=new Date,a.lastModified=
b.getTime(),a.lastModifiedDate=b,a.name=d.name,a.name&&a.type!==d.type&&(a.name=a.name.replace(S,v(a.type)))):a=d;this.result=a;e.success&&e.success.call(this,a)}},{key:"fail",value:function(a){var b=this.options;if(b.error)b.error.call(this,a);else throw a;}},{key:"abort",value:function(){this.aborted||(this.aborted=!0,this.reader?this.reader.abort():this.image.complete?this.fail(Error("The compression process has been aborted.")):(this.image.onload=null,this.image.onabort()))}}],[{key:"noConflict",
value:function(){window.Compressor=T;return a}},{key:"setDefaults",value:function(a){e(J,a)}}]);return a}()})},{}]},{},[1])(1)});
