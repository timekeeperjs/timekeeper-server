import api from "./api";

export interface RemoteResponse {
  remoteName: string;
  remoteURL: string;
  version: string;
}

// interface ErrorResponse {
//   error: string;
// }

interface PushRemoteRequest {
  baseUrl: string;
  remoteName: string;
  version: string;
}

export const getAllRemotes = async (
  remoteName?: string,
): Promise<RemoteResponse[]> => {
  const response = await api.get<RemoteResponse[]>("/dashboard", {
    params: { remoteName },
  });
  return response.data;
};

export const getRemoteByNameAndVersion = async (
  remoteName: string,
  version?: string,
): Promise<RemoteResponse> => {
  const response = await api.get<RemoteResponse>("/get-remote", {
    params: { remoteName, version },
  });
  return response.data;
};

export const healthCheck = async (): Promise<{ status: string }> => {
  const response = await api.get<{ status: string }>("/health-check");
  return response.data;
};

export const pushRemote = async (
  remote: PushRemoteRequest,
): Promise<RemoteResponse> => {
  const response = await api.post<RemoteResponse>("/push-remote", remote);
  return response.data;
};

export const getAllUniqueRemoteNames = async (): Promise<string[]> => {
  const response = await api.get<string[]>("/remote-names");
  return response.data;
};
