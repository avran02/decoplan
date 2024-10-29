// types.ts
export interface Attachment {
  id: string;
  url: string;
}

export interface MessageContent {
  text: string;
  attachments?: Attachment[];
}

export interface Message {
  id: string;
  sender: string;
  content: MessageContent;
  timestamp: string;
}

export interface AskMessagesDto {
  chatId: string;
  limit?: number;
  offset?: number;
}

export interface DeleteMessageDto {
  chatId: string;
  messageId: number;
}

export interface NewMessageDto {
  chatId: string;
  content: MessageContent;
  timestamp: string;
}
