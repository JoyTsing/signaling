"use strict";

let pushBtn = document.getElementById("pushBtn");

let uid = $("#uid").val();
let streamName = $("#streamName").val();
let audio = $("#audio").val();
let video = $("#video").val();

pushBtn.addEventListener("click", function () {
  $.post(
    "/signaling/push",
    {
      uid: uid,
      streamName: streamName,
      audio: audio,
      video: video,
    },
    function (data, textStatus) {}
  );
});
