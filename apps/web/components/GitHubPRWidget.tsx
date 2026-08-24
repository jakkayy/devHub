'use client';

import React from 'react';
import { useQuery } from '@tanstack/react-query';
import { fetchGitHubPullRequests } from '@/services/api';
import { ExternalLink, GitPullRequest } from 'lucide-react';

export default function GitHubPRWidget() {
  const { data: prs, isLoading, isError } = useQuery({
    queryKey: ['github-prs'],
    queryFn: fetchGitHubPullRequests,
  });

  return (
    <div className="bg-zinc-900/90 border border-zinc-800 rounded-xl p-5 shadow-lg">
      <div className="flex items-center justify-between mb-4">
        <div className="flex items-center gap-2">
          <GitPullRequest className="w-5 h-5 text-purple-400" />
          <h2 className="text-lg font-semibold text-zinc-100">Review Queue (PRs)</h2>
        </div>
        <span className="text-xs text-zinc-400">GitHub Activity</span>
      </div>

      {isLoading && <div className="py-8 text-center text-sm text-zinc-500">Loading PRs...</div>}
      {isError && <div className="py-8 text-center text-sm text-rose-400">Failed to load PRs</div>}

      {prs && prs.length === 0 && (
        <div className="py-8 text-center text-sm text-zinc-500">No active Pull Requests</div>
      )}

      {prs && prs.length > 0 && (
        <div className="space-y-3">
          {prs.map((pr) => (
            <div key={pr.id} className="flex items-center justify-between p-3 rounded-lg bg-zinc-950/60 border border-zinc-800/80 hover:border-zinc-700 transition">
              <div>
                <div className="flex items-center gap-2">
                  <span className="text-xs font-mono font-medium text-purple-400">#{pr.number}</span>
                  <p className="text-sm font-medium text-zinc-200">{pr.title}</p>
                </div>
                <p className="text-xs text-zinc-500 mt-1">Author: @{pr.author}</p>
              </div>
              <a
                href={pr.html_url}
                target="_blank"
                rel="noreferrer"
                className="flex items-center gap-1 text-xs text-purple-400 hover:text-purple-300 font-medium transition"
              >
                Review <ExternalLink className="w-3 h-3" />
              </a>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
