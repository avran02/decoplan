// useChatWebSocket.ts
import { WS_URL } from '@/constants/constants'
import { AskMessagesDto, DeleteMessageDto, Message, NewMessageDto } from '@/types/chat.types'
import { useCallback, useEffect, useState } from 'react'
import useWebSocket, { ReadyState } from 'react-use-websocket'

const useChatWebSocket = (token: string) => {
  const [messages, setMessages] = useState<Message[]>([]);
  const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket(WS_URL, {
    onOpen: () => console.log('WebSocket Connection Opened'),
    onClose: () => console.log('WebSocket Connection Closed'),
    queryParams: { token },
    shouldReconnect: () => true,
  });

  // Обновление сообщений при получении нового сообщения
  useEffect(() => {
    if (lastJsonMessage) {
      setMessages((prev) => [...prev, lastJsonMessage as Message]);
    }
  }, [lastJsonMessage]);

  // Функция для подключения
  const connect = useCallback(() => {
    if (readyState !== ReadyState.OPEN) {
      console.log('Connecting WebSocket...');
    }
  }, [readyState]);

  // Функция для отключения
  const disconnect = useCallback(() => {
    if (readyState === ReadyState.OPEN) {
      sendJsonMessage({ type: 'disconnect' });
    }
  }, [readyState, sendJsonMessage]);

  // Отправка нового сообщения
  const sendMessage = useCallback((messageData: NewMessageDto) => {
    sendJsonMessage({ type: 'new_message', data: messageData });
  }, [sendJsonMessage]);

  // Получение списка сообщений
  const fetchMessages = useCallback((data: AskMessagesDto) => {
    sendJsonMessage({ type: 'fetch_messages', data });
  }, [sendJsonMessage]);

  // Удаление сообщения
  const deleteMessage = useCallback((data: DeleteMessageDto) => {
    sendJsonMessage({ type: 'delete_message', data });
  }, [sendJsonMessage]);

  return {
    messages,
    connect,
    disconnect,
    sendMessage,
    fetchMessages,
    deleteMessage,
    readyState,
  };
};

export default useChatWebSocket;
