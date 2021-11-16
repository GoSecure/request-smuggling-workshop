package net.gosecure.websocket.config;

import java.io.IOException;
import java.util.Random;

import org.json.JSONObject;
import org.springframework.stereotype.Component;
import org.springframework.web.socket.TextMessage;
import org.springframework.web.socket.WebSocketSession;
import org.springframework.web.socket.handler.TextWebSocketHandler;

@Component
public class SocketTextHandler extends TextWebSocketHandler {

	private String[] CANNED_ANSWERS = {"How may we help you?",
			"Was there anything else?",
			"I'm sorry I have no idea..",
			"Do you have more details?",
			"That's interesting..",
			"I'm not sure I understand.. Can you rephrase it?",
			"Just give me a moment.",
			"Is it something recent?",
			"That's not very clear.",
			"I need more details.",
			"And..."};


	@Override
	public void handleTextMessage(WebSocketSession session, TextMessage message)
			throws InterruptedException, IOException {

		String payload = message.getPayload();
		JSONObject jsonObject = new JSONObject(payload);
		session.sendMessage(new TextMessage("<b>User</b>: " + jsonObject.get("user")));
		Thread.sleep(2000);

		String response = CANNED_ANSWERS[new Random().nextInt(CANNED_ANSWERS.length)];
		session.sendMessage(new TextMessage("<b>Support:</b>: "+response));

	}

}