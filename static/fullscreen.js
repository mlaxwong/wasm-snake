(function () {
  function getFullscreenElement() {
    let e = document;
    return document.fullscreenElement   //standard property
    || document.webkitFullscreenElement //safari/opera support
    || document.mozFullscreenElement    //firefox support
    || document.msFullscreenElement;    //ie/edge support
  }

  function toggleFullscreen() {
    if(getFullscreenElement()) {
      document.exitFullscreen();
    } else {
      // document.documentElement.requestFullscreen().catch(console.log);
      document.getElementById("screen").requestFullscreen()
    }
  }

  button = document.getElementById("fullscreen")
  button.addEventListener('click', () => {
    toggleFullscreen();
  });

  // document.documentElement.requestFullscreen().catch(console.log);

  // document.getElementById("screen").requestFullscreen()
})()