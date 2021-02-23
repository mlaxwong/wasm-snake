(function () {
  let pageWidth = window.innerWidth || document.body.clientWidth;
  let treshold = Math.max(1,Math.floor(0.01 * (pageWidth)));
  let touchstartX = 0;
  let touchstartY = 0;
  let touchendX = 0;
  let touchendY = 0;

  const limit = Math.tan(45 * 1.5 / 180 * Math.PI);
  const gestureZone = window

  gestureZone.addEventListener('touchstart', function(event) {
      touchstartX = event.changedTouches[0].screenX;
      touchstartY = event.changedTouches[0].screenY;
  }, false);

  gestureZone.addEventListener('touchend', function(event) {
      touchendX = event.changedTouches[0].screenX;
      touchendY = event.changedTouches[0].screenY;
      handleGesture(event);
  }, false);

  function handleGesture(e) {
      let x = touchendX - touchstartX;
      let y = touchendY - touchstartY;
      let xy = Math.abs(x / y);
      let yx = Math.abs(y / x);
      if (Math.abs(x) > treshold || Math.abs(y) > treshold) {
          if (yx <= limit) {
              if (x < 0) {
                  // console.log("left");
                  gestureZone.dispatchEvent(new CustomEvent('swipe', {
                    bubbles: true,
                    detail: 3,
                  }));
              } else {
                  // console.log("right");
                  gestureZone.dispatchEvent(new CustomEvent('swipe', {
                    bubbles: true,
                    detail: 4,
                  }));
              }
          }
          if (xy <= limit) {
              if (y < 0) {
                  // console.log("top");
                  gestureZone.dispatchEvent(new CustomEvent('swipe', {
                    bubbles: true,
                    detail: 1,
                  }));;
              } else {
                  // console.log("bottom");
                  gestureZone.dispatchEvent(new CustomEvent('swipe', {
                    bubbles: true,
                    detail: 2,
                  }));
              }
          }
      } else {
          // console.log("tap");
          gestureZone.dispatchEvent(new CustomEvent('tap'));
      }
  }

  // function getFullscreenElement() {
  //     return document.fullscreenElement   //standard property
  //     || document.webkitFullscreenElement //safari/opera support
  //     || document.mozFullscreenElement    //firefox support
  //     || document.msFullscreenElement;    //ie/edge support
  // }

  // function toggleFullscreen() {
  //     if(getFullscreenElement()) {
  //       document.exitFullscreen();
  //     }else {
  //   document.documentElement.requestFullscreen().catch(console.log);
  //     }
  // }

  // game = document.getElementById("game")
  // document.addEventListener('dblclick', () => {
  //     toggleFullscreen();
  // });

  // window.addEventListener('swipe', function(event) {
  //   console.log(event)
  // })
})()