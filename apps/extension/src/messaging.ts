export type ExtensionMessageType =
  | 'FETCH_TASKS'
  | 'FETCH_CONTRACTS'
  | 'LINK_CODE'
  | 'OPEN_DEEP_LINK'
  | 'NOTIFY';

export interface WebviewMessage {
  command: ExtensionMessageType;
  payload?: any;
}

export interface ExtensionResponse {
  type: string;
  data?: any;
  error?: string;
}
