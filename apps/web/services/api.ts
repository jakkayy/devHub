export interface Task {
  id: string;
  title: string;
  assignee: string;
  status: string;
  source: string;
  external_url: string;
  updated_at: string;
}

export interface GitHubPullRequest {
  id: number;
  number: number;
  title: string;
  state: string;
  author: string;
  html_url: string;
  draft: boolean;
  created_at: string;
  updated_at: string;
}

export interface APIContract {
  id: string;
  method: string;
  path: string;
  summary: string;
  spec_version: string;
}

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1';

export async function fetchTasks(): Promise<Task[]> {
  const res = await fetch(`${API_BASE_URL}/tasks`);
  if (!res.ok) throw new Error('Failed to fetch tasks');
  const json = await res.json();
  return json.data || [];
}

export async function fetchGitHubPullRequests(): Promise<GitHubPullRequest[]> {
  const res = await fetch(`${API_BASE_URL}/github/pulls`);
  if (!res.ok) throw new Error('Failed to fetch GitHub PRs');
  const json = await res.json();
  return json.data || [];
}

export async function fetchAPIContracts(): Promise<APIContract[]> {
  const res = await fetch(`${API_BASE_URL}/contracts`);
  if (!res.ok) throw new Error('Failed to fetch API contracts');
  const json = await res.json();
  return json.data || [];
}
