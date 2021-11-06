function timeformat(timestamp) {
    return new Date(parseInt(timestamp) * 1000).toLocaleString().replace(/:\d{1,2}$/,' '); 
}

function makeNotification(event) {
  var popNotice = function(event) {
    if (Notification.permission == "granted") {
        var notification = new Notification(event.Title, {
            body: event.Description,
            icon: 'https://avatar.csdn.net/0/8/F/3_marksinoberg.jpg'
        });
        
        notification.onclick = function() {
            // alert(event.Description)
            notification.close();    
        };
    }    
};

  if(Notification.permission == "granted") {
    popNotice(event);   
  }else if(Notification.permission != "denied") {
    Notification.requestPermission(function(permission){
        popNotice(event);
    });
  }
}

function appendEvent(events) {
  for(index=0; index<events.length; index++) {
    var childstr = "<li><div title='"+events[index].Description+"'>"+"<strong><blob><code>"+timeformat(events[index].Tiptime)+"</code></blob></strong><br>"+events[index].Title+"</div><br></li>";
    console.log(childstr);
    $("#bucket").prepend(childstr);
  }
}


function getEvents(all=false) {
  console.log("ready to get events");
  var timestamp = (Date.parse(new Date()))/1000;
    if(all == true) {
      var starttime = 0;
      var endtime = (Date.parse(new Date()))/1000 + 86400 * 90;
    }else{
      var starttime = (Date.parse(new Date()))/1000 - 86400*3;
      var endtime = (Date.parse(new Date()))/1000 + 86400*4;
    }
    $.ajax({
      type:"get",
      url:"/getevent",
      data:{starttime:starttime, endtime: endtime},
      dataType:"json",
      cache:false,
      async:true,
      success:function(result){
          console.info(result['ret'].length);
          var events = result['ret'];
          if(events && events.length > 0) {
              appendEvent(events);
              // makeNotification(events[0]);
          }else{
              $("#namedlist").html("temporarily set to null");
          }
      }
  });
}

function getEventForNotification() {
  console.log("ready to get events");
  var timestamp = Date.parse(new Date())/1000;
  var starttime = (Date.parse(new Date()))/1000 - 86400*3;
  var endtime = (Date.parse(new Date()))/1000 + 86400*4;
    $.ajax({
      type:"get",
      url:"/getevent",
      data:{starttime:starttime, endtime: endtime},
      dataType:"json",
      cache:false,
      async:true,
      success:function(result){
          console.info(result['ret'].length);
          var events = result['ret'];
          if(events && events.length > 0) {
              // notification within 1m
              for(index=0; index<events.length; index++) {
                if(events[index].Tiptime > timestamp && events[index].Tiptime <= timestamp + 60) {
                  makeNotification(events[index]);
                }
              }
          }else{
            $("#namedlist").html("temporarily set to null");
          }
      }
  });
}

function addEvent(title, description) {
  $.ajax({
    type: "get",
    url: "/addevent",
    data: {title: title, description:description},
    dataType: "json",
    cache: false,
    async: true,
    success: function(result) {
      console.log(result);
      $("#btn_addevent").html("添加");
      $("#input_title").attr("value", "");
      $("#input_description").attr("value", "")
    },
    error: function(error) {
      console.log(error)
    }
  });
}
$("#btn_generate").click(function(){
  var title = $("#input_title").val();
  var description = $("#input_description").val();
  console.log(title, description);
  if(title!="" && description.length>11) {
    addEvent(title, description);
    location.reload();
  }else{
    alert("description too short");
  }
});

$(document).ready(function(){
  getEvents(false);
});
setInterval(getEventForNotification, 60*1000);
// setTimeout(makeNotification, 2000);
