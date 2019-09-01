define('imageUpload', ['ajax', 'accto', 'Compressor'], function(ajax, accto, Compressor){

  function isImage(file, allowExts){
  	if(!file) return false;
  	var arr = file.type.split("/");
  	if(arr[0]!="image") return false;

  	return true;
  }

  return {
    template:
    `<div class="gno-image-upload">
      <p>----------上传图片:----------</p>
        <input ref="input-file" type="file" accept="image/*">
        <label>
          <button @click="_upload">上传</button>
        </label>
      <div ref="imgdiv"></div>
    </div>`
    ,
    props : {
      quality: Number
    },
    methods: {
      _upload: function(e){
				var files = this.$refs['input-file'].files;
				if(files.length==0) return;
				var imgdiv = this.$refs['imgdiv'], img = document.createElement("img");
        var file = files[0];
        var str = file.name +  file.size + file.lastModified + file.type+navigator.userAgent + window.location.hostname;

        if(this.quality===0) this.quality = 0.6;

        new Compressor(file, {
          quality: this.quality,
          maxWidth: 1024*2,
          maxHeight: 1024*2,
          success(blob) {
  					img.src = URL.createObjectURL(blob);
            imgdiv.appendChild(img);
            ajax.NewUpload('/upload/app').upload(blob, accto(str));
          },
          error(err) {
            console.log(err.message);
          },
        });

      }

    }
  };
})
