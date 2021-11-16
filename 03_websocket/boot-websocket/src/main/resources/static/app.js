//var WS_URL = "ws://localhost:8081/chat";

var isConnected = false

var ws;
function setConnected(connected) {
    isConnected = connected;
	$("#connect").prop("disabled", connected);
	$("#disconnect").prop("disabled", !connected);
}

function connect() {
    WS_URL = "ws://"+location.host+"/chat";

	ws = new WebSocket(WS_URL);
	ws.onmessage = function(data) {
		helloWorld(data.data);
	}
	setConnected(true);
}

function disconnect() {
	if (ws != null) {
		ws.close();
	}
	setConnected(false);
	console.log("Websocket is in disconnected state");
}

function sendData() {
    var wait = !isConnected;
    if(!isConnected) {
        connect();
    }

    setTimeout(function(){
        var data = JSON.stringify({
    		'user' : $("#user").val()
    	})
    	ws.send(data);

     }, wait? 2000:0);


}

function helloWorld(message) {
	$("#helloworldmessage").append("<tr><td> " + message + "</td></tr>");
}

$(function() {
	$("form").on('submit', function(e) {
		e.preventDefault();
	});
	$("#connect").click(function() {
		connect();
	});
	$("#disconnect").click(function() {
		disconnect();
	});
	$("#send").click(function() {
		sendData();
	});
});
