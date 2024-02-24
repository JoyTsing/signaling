"use strict";

var pushBtn = document.getElementById("pushBtn");

var uid = $("#uid").val();
var streamName = $("#streamName").val();
var audio = $("#audio").val();
var video = $("#video").val();

pushBtn.addEventListener("click", function () {
  $.post(
    "/signaling/push",
    {
      uid: uid,
      streamName: streamName,
      audio: audio,
      video: video,
    },
    function (data, textStatus) {
      console.log("push response: " + JSON.stringify(data));
      if ("success" == textStatus && 0 == data.errNo) {
        $("#tips").html("<font color='blue'>推流成功</font>");
      } else {
        $("#tips").html("<font color='red'>推流失败</font>");
      }
    },
    "json"
  );
});
