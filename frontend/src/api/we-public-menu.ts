import axios from 'axios';

export interface GetOAMenuTreeReply {
  button: any;
  matchrule: any;
}

export function getOAMenuTree() {
  return axios.get<GetOAMenuTreeReply>(`/we-public-menu/menus-tree`);
}

export interface SyncOAMenuRequest {
  button: any;
}

export interface SyncOAMenuReply {
  success: boolean;
  data: any;
}

export function syncOAMenu(request: SyncOAMenuRequest) {
  return axios.post<SyncOAMenuReply>(`/we-public-menu/menus`, request);
}
