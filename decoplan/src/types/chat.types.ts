// types.ts
export interface IAttachment {
  id: string;
  url: string;
}

export interface IMessageContent {
  text: string;
  attachments?: IAttachment[];
}

export interface IMessage {
  id: string;
  sender: string;
  content: IMessageContent;
  timestamp: string;
}

export interface IAskMessagesDto {
  chatId: string;
  limit?: number;
  offset?: number;
}

export interface IDeleteMessageDto {
  chatId: string;
  messageId: number;
}
export interface INewMessageDto {
  chatId: string;
  content: IMessageContent;
  timestamp: string;
}


export enum UserMessages {
  USER_GET_MESSAGES = 0,
  USER_SEND_MESSAGE = 1,
  USER_DELETE_MESSAGE = 2,
}