'use client';

import React from 'react';
import { ExternalLink, FileCode, GitBranch, Layers, Terminal } from 'lucide-react';

export interface NodeLinkMapping {
  nodeId: string;
  nodeName: string;
  taskId: string;
  apiContractId: string;
  filePath: string;
  lineNumber: number;
  vscodeUri: string;
}

interface MiroNodeLinkPanelProps {
  mappings?: NodeLinkMapping[];
}

export default function MiroNodeLinkPanel({ mappings }: MiroNodeLinkPanelProps) {
  // Default mock mappings if not provided
  const sampleMappings: NodeLinkMapping[] = mappings || [
    {
      nodeId: 'node-1',
      nodeName: 'User Auth & Task Handler',
      taskId: 'TASK-101',
      apiContractId: 'GET-api-v1-tasks',
      filePath: '/home/naeiger/project-personal/platform-engineer/devHub/services/backend-api/cmd/api/main.go',
      lineNumber: 25,
      vscodeUri: 'vscode://file//home/naeiger/project-personal/platform-engineer/devHub/services/backend-api/cmd/api/main.go:25',
    },
    {
      nodeId: 'node-2',
      nodeName: 'Miro Diagram Parser Engine',
      taskId: 'TASK-103',
      apiContractId: 'GET-api-v1-contracts',
      filePath: '/home/naeiger/project-personal/platform-engineer/devHub/apps/web/components/MiroWidget.tsx',
      lineNumber: 1,
      vscodeUri: 'vscode://file//home/naeiger/project-personal/platform-engineer/devHub/apps/web/components/MiroWidget.tsx:1',
    },
  ];

  return (
    <div className="bg-zinc-900/90 border border-zinc-800 rounded-xl p-5 shadow-lg space-y-4">
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-2">
          <Layers className="w-5 h-5 text-amber-400" />
          <h3 className="text-base font-semibold text-zinc-100">Node-to-Code Links</h3>
        </div>
        <span className="text-xs text-zinc-400">Context Linker Engine</span>
      </div>

      <div className="space-y-3">
        {sampleMappings.map((item) => (
          <div
            key={item.nodeId}
            className="p-3 rounded-lg bg-zinc-950/60 border border-zinc-800 hover:border-amber-500/40 transition space-y-2"
          >
            <div className="flex items-center justify-between">
              <span className="text-xs font-semibold text-amber-400">{item.nodeName}</span>
              <span className="text-[10px] font-mono px-2 py-0.5 rounded bg-zinc-800 text-zinc-300">
                {item.taskId}
              </span>
            </div>

            <div className="flex flex-col gap-1 text-xs text-zinc-400">
              <div className="flex items-center gap-1.5 font-mono text-[11px] text-zinc-300">
                <FileCode className="w-3.5 h-3.5 text-sky-400" />
                <span className="truncate">{item.filePath.split('/').pop()}:{item.lineNumber}</span>
              </div>

              {item.apiContractId && (
                <div className="flex items-center gap-1.5 text-[11px] text-zinc-400">
                  <GitBranch className="w-3.5 h-3.5 text-purple-400" />
                  <span>API: {item.apiContractId}</span>
                </div>
              )}
            </div>

            <div className="pt-1 flex items-center justify-end">
              <a
                href={item.vscodeUri}
                className="inline-flex items-center gap-1 text-xs text-emerald-400 hover:text-emerald-300 font-medium bg-emerald-950/60 border border-emerald-800/40 px-2.5 py-1 rounded-md transition"
              >
                <Terminal className="w-3 h-3" /> Open in VS Code <ExternalLink className="w-3 h-3" />
              </a>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
