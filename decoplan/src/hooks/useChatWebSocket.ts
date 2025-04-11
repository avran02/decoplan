import { WS_URL } from '@/constants/constants'
import { IAskMessagesDto, IDeleteMessageDto, IMessage, INewMessageDto, UserMessages } from '@/types/chat.types'
import { useCallback, useState } from 'react'
import useWebSocket, { ReadyState } from 'react-use-websocket'

const useChatWebSocket = (token: string) => {
  const [messages, setMessages] = useState<IMessage[]>([]);
  const { sendJsonMessage, readyState } = useWebSocket(WS_URL, {
    onOpen: () => console.log('WebSocket Connection Opened'),
    onClose: () => console.log('WebSocket Connection Closed'),
    onMessage: (event) => { 
      if (event.data.length) {
        const WSMessage = JSON.parse(event.data);

        if (WSMessage.payload) {
          setMessages((prevMessages) => {
            const newMessages = WSMessage.payload as IMessage[];
            const existingIds = new Set(prevMessages.map(msg => msg.id));
            const filteredMessages = newMessages.filter(msg => !existingIds.has(msg.id));
            
            return [...prevMessages, ...filteredMessages];
          });
        }
      }
    },
    queryParams: { token },
    shouldReconnect: () => true,
  });

  const connect = useCallback(() => {
    if (readyState !== ReadyState.OPEN) {
      console.log('Connecting WebSocket...');
    }
  }, [readyState]);

  const disconnect = useCallback(() => {
    if (readyState === ReadyState.OPEN) {
      sendJsonMessage({ type: 'disconnect' });
    }
  }, [readyState, sendJsonMessage]);

  const sendMessage = (messageData: INewMessageDto) => {
    sendJsonMessage({ act: UserMessages.USER_SEND_MESSAGE, payload: messageData });
  };

  const fetchMessages = (data: IAskMessagesDto) => {
    sendJsonMessage({ act: UserMessages.USER_GET_MESSAGES, payload: data });
  };

  const deleteMessage = (data: IDeleteMessageDto) => {
    sendJsonMessage({ act: UserMessages.USER_DELETE_MESSAGE, payload: data });
    // Locally remove the message after deletion
    setMessages((prevMessages) => prevMessages.filter(msg => msg.id !== data.messageId));
  };

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
