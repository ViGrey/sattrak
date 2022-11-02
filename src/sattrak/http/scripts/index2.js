var homeLocation = document.getElementById("location");
var homeLatitude = document.getElementById("home_latitude");
var homeLongitude = document.getElementById("home_longitude");
var startTime = document.getElementById("start_time");
var startTimeNow = document.getElementById("start_time_now");

function setStartTimeToNow() {
  var timeNow = new Date();

  console.log(timeNow.getUTCDate());
  console.log(timeNow.getUTCMonth());
  startTime.value = timeNow.toISOString().slice(0,16);
}

function getLocation() {
  if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(handleCurrentPosition)
  }
}

function handleCurrentPosition(position) {
  home_latitude.value = position.coords.latitude;
  home_longitude.value = position.coords.longitude;
}

startTimeNow.addEventListener("change", () => {
  if (startTimeNow.checked) {
    startTime.setAttribute("disabled", "");
    setStartTimeToNow();
  } else {
    startTime.removeAttribute("disabled");
  }
});

homeLocation.addEventListener("click", () => {
  getLocation();
});
