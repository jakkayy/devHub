'use client';

import React from 'react';
import { useQuery } from '@tanstack/react-query';
import { fetchAPIContracts } from '../services/api';
import { Code2, FileCode2 } from 'lucide-react';

export default function APIContractWidget() {
  const { data: contracts, isLoading, isError } = useQuery({
    queryKey: ['api-contracts'],
    queryFn: fetchAPIContracts,
  });

  const getMethodBadge = (method: string) => {
    switch (method.toUpperCase()) {
      case 'GET':
        return <span className="text-xs font-mono font-bold px-2 py-0.5 rounded bg-blue-950 text-blue-400 border border-blue-800/50">GET</span>;
      case 'POST':
        return <span className="text-xs font-mono font-bold px-2 py-0.5 rounded bg-emerald-950 text-emerald-400 border border-emerald-800/50">POST</span>;
      case 'PUT':
        return <span className="text-xs font-mono font-bold px-2 py-0.5 rounded bg-amber-950 text-amber-400 border border-amber-800/50">PUT</span>;
      case 'DELETE':
        return <span className="text-xs font-mono font-bold px-2 py-0.5 rounded bg-rose-950 text-rose-400 border border-rose-800/50">DELETE</span>;
      default:
        return <span className="text-xs font-mono font-bold px-2 py-0.5 rounded bg-zinc-800 text-zinc-300">{method}</span>;
    }
  };

  return (
    <div className="bg-zinc-900/90 border border-zinc-800 rounded-xl p-5 shadow-lg">
      <div className="flex items-center justify-between mb-4">
        <div className="flex items-center gap-2">
          <FileCode2 className="w-5 h-5 text-sky-400" />
          <h2 className="text-lg font-semibold text-zinc-100">API Contract Specs</h2>
        </div>
        <span className="text-xs text-zinc-400">OpenAPI 3.0</span>
      </div>

      {isLoading && <div className="py-8 text-center text-sm text-zinc-500">Loading specs...</div>}
      {isError && <div className="py-8 text-center text-sm text-rose-400">Failed to load specs</div>}

      {contracts && contracts.length === 0 && (
        <div className="py-8 text-center text-sm text-zinc-500">No API specs available</div>
      )}

      {contracts && contracts.length > 0 && (
        <div className="space-y-3">
          {contracts.map((contract) => (
            <div key={contract.id} className="flex items-center justify-between p-3 rounded-lg bg-zinc-950/60 border border-zinc-800/80 hover:border-zinc-700 transition">
              <div className="flex items-center gap-3">
                {getMethodBadge(contract.method)}
                <div>
                  <p className="text-sm font-mono text-zinc-200">{contract.path}</p>
                  <p className="text-xs text-zinc-500 mt-0.5">{contract.summary}</p>
                </div>
              </div>
              <Code2 className="w-4 h-4 text-zinc-600" />
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
