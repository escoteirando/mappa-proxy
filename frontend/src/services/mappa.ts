import { api } from 'src/boot/axios';

export interface ILogin {
  id: string;
  ttl: number;
  created: Date;
  userId: number;
}

export interface IMappaServerHC {
  status: string;
  status_code: number;
}
export interface IProxyHC {
  mappa_server: IMappaServerHC;
  status: string;
}

export interface ILoginContext {
  cId: number;
  mId: number;
}
export interface IChatPostResponse {
  message: string;
  error: unknown;
}

export const EmptyLoginContext: ILoginContext = {
  cId: 0,
  mId: 0,
};

export function DoLogin(username: string, password: string): Promise<ILogin> {
  return new Promise((resolve, reject) => {
    const body = {
      type: 'LOGIN_REQUEST',
      username: username,
      password: password,
    };
    api
      .post('/mappa/login', body)
      .then((response) => {
        resolve(response.data);
      })
      .catch((error) => {
        reject(error);
      });
  });
}

export function SendAuthToChat(
  cId: number,
  mId: number,
  msg: string
): Promise<IChatPostResponse> {
  return new Promise((resolve, reject) => {
    const body = {
      cId: parseInt(`${cId}`),
      mId: parseInt(`${mId}`),
      msg: msg,
    };
    console.log('Sending auth', body);  // TODO: Remover console logs
    api
      .post('/tg/pub', body)
      .then((response) => {
        console.log('SendAuthToChat', response);
        resolve(response as unknown as IChatPostResponse);
      })
      .catch((error) => {
        console.error('SendAuthToChat', error);
        reject(error as IChatPostResponse);
      });
  });
}

export function DoHC(): Promise<IProxyHC> {
  return new Promise((resolve, reject) => {
    api
      .get('/hc')
      .then((response) => resolve(response.data as IProxyHC))
      .catch((error) => {
        reject(error);
      });
  });
}

export function ParseContext(context: string): ILoginContext {
  try {
    const j = atob(context);
    const jsonContext: ILoginContext = JSON.parse(j) as ILoginContext;
    if (!jsonContext.cId) {
      throw new Error(`Invalid context data: ${JSON.stringify(jsonContext)}`);
    }
    return {
      cId: jsonContext.cId,
      mId: jsonContext.mId,
    };
  } catch (error) {
    console.error('ParseContent', error);
    return EmptyLoginContext;
  }
}
