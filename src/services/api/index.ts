import { AxiosInstance, AxiosResponse } from 'axios';

import { sleep } from '@/utils';
import logger from '@/services/logger';
import { loggers } from 'winston';

export * from '@/services/api/github';

export type APIType = 'github' | 'gitlab' | 'bitbucket' | string;
export type AuthMethod = 'token' | 'password';
export type AuthState = 'active' | 'inactive';

export interface IAPI {
  readonly name: APIType,
  auth: () => Promise<any>,
  // eslint-disable-next-line no-unused-vars
  get: (url: string) => Promise<any>,
}

export interface IPollConfig {
  interval: number,
  timeout: number
}

export async function postPoll(
  config: IPollConfig,
  instance: AxiosInstance,
  url: string,
  data?: any,
): Promise<AxiosResponse> {
  let pollResponse = await instance.post(url, data);
  const startTime = Date.now();
  while (pollResponse.data.access_token === undefined
    && startTime - Date.now() < config.timeout) {
    // eslint-disable-next-line no-loop-func
    // eslint-disable-next-line no-await-in-loop
    pollResponse = await instance.post(url, data);
    // eslint-disable-next-line no-await-in-loop
    await sleep(pollResponse.data.interval * 1000);
  }
  return pollResponse;
}

export interface IService {
  readonly name: APIType,
  readonly authMethod: AuthMethod | Array<AuthMethod>,
  readonly apiUrl: string
}

export interface IAPIOptions {
  service: string | APIType,
}
