// useChatWebSocket.ts
import { WS_URL } from '@/constants/constants'
import { IAskMessagesDto, IDeleteMessageDto, IMessage, INewMessageDto, UserMessages } from '@/types/chat.types'
import { useCallback, useEffect, useState } from 'react'
import useWebSocket, { ReadyState } from 'react-use-websocket'

const useChatWebSocket = (token: string) => {
  const [messages, setMessages] = useState<IMessage[]>([]);
  const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket(WS_URL, {
    onOpen: () => console.log('WebSocket Connection Opened'),
    onClose: () => console.log('WebSocket Connection Closed'),
    queryParams: { token },
    shouldReconnect: () => true,
  });

  // Обновление сообщений при получении нового сообщения
  useEffect(() => {
    if (lastJsonMessage) {
      setMessages((prev) => [...prev, lastJsonMessage as IMessage]);
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
  const sendMessage = useCallback((messageData: INewMessageDto) => {
    sendJsonMessage({ act: UserMessages.USER_SEND_MESSAGE, payload: messageData });
  }, [sendJsonMessage]);

  // Получение списка сообщений
  const fetchMessages = useCallback((data: IAskMessagesDto) => {
    sendJsonMessage({ act: UserMessages.USER_GET_MESSAGES, payload: data });
  }, [sendJsonMessage]);

  // Удаление сообщения
  const deleteMessage = useCallback((data: IDeleteMessageDto) => {
    sendJsonMessage({ act: UserMessages.USER_DELETE_MESSAGE, payload: data });
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
