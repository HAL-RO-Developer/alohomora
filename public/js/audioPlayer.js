(function (exports) {

	function AudioPlayer(){
		this.audioCtx = new (window.AudioContext || window.webkitAudioContext)();
		this.source = null;
		this.isSomeMusicPlaying = false;
	}

	AudioPlayer.prototype.getData = function(filename){
		var self = this;

		this.source = this.audioCtx.createBufferSource()
		var request = new XMLHttpRequest()

		request.open("GET",filename,true)
		request.responseType = "arraybuffer";

		request.onload = function(){
			var audioData = request.response;

			self.audioCtx.decodeAudioData(audioData,function(buffer){
				self.source.buffer = buffer;

				self.source.connect(self.audioCtx.destination)
			})
		}

		request.send()
	}

	AudioPlayer.prototype.play = function(filename){
		this.getData(filename)
		this.isSomeMusicPlaying = true;
		this.source.start(0);
	}

	AudioPlayer.prototype.stop = function(){
		this.isSomeMusicPlaying = false;
		this.source.stop(0)
	}

	AudioPlayer.prototype.isSomeMusicPlaying = function(){
		return this.isSomeMusicPlaying;
	}

  exports.AudioPlayer = AudioPlayer;
}(window));
