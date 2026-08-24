'use client';

import React from 'react';
import { useQuery } from '@tanstack/react-query';
import { fetchTasks } from '@/services/api';
import { CheckCircle2, Clock, ListTodo, Sheet } from 'lucide-react';

export default function TaskWidget() {
  const { data: tasks, isLoading, isError } = useQuery({
    queryKey: ['tasks'],
    queryFn: fetchTasks,
  });

  const getStatusBadge = (status: string) => {
    switch (status.toLowerCase()) {
      case 'done':
        return <span className="inline-flex items-center gap-1 text-xs px-2 py-0.5 rounded-full bg-emerald-950/80 text-emerald-400 border border-emerald-800/50"><CheckCircle2 className="w-3 h-3" /> Done</span>;
      case 'in progress':
        return <span className="inline-flex items-center gap-1 text-xs px-2 py-0.5 rounded-full bg-amber-950/80 text-amber-400 border border-amber-800/50"><Clock className="w-3 h-3" /> In Progress</span>;
      default:
        return <span className="inline-flex items-center gap-1 text-xs px-2 py-0.5 rounded-full bg-zinc-800 text-zinc-300 border border-zinc-700"><ListTodo className="w-3 h-3" /> {status}</span>;
    }
  };

  return (
    <div className="bg-zinc-900/90 border border-zinc-800 rounded-xl p-5 shadow-lg">
      <div className="flex items-center justify-between mb-4">
        <div className="flex items-center gap-2">
          <Sheet className="w-5 h-5 text-emerald-400" />
          <h2 className="text-lg font-semibold text-zinc-100">Tasks Tracking</h2>
        </div>
        <span className="text-xs text-zinc-400">Google Sheets Sync</span>
      </div>

      {isLoading && <div className="py-8 text-center text-sm text-zinc-500">Loading tasks...</div>}
      {isError && <div className="py-8 text-center text-sm text-rose-400">Failed to load tasks</div>}

      {tasks && tasks.length === 0 && (
        <div className="py-8 text-center text-sm text-zinc-500">No tasks found</div>
      )}

      {tasks && tasks.length > 0 && (
        <div className="space-y-3">
          {tasks.map((task) => (
            <div key={task.id} className="flex items-center justify-between p-3 rounded-lg bg-zinc-950/60 border border-zinc-800/80 hover:border-zinc-700 transition">
              <div>
                <div className="flex items-center gap-2">
                  <span className="text-xs font-mono font-medium text-emerald-400">{task.id}</span>
                  <p className="text-sm font-medium text-zinc-200">{task.title}</p>
                </div>
                <p className="text-xs text-zinc-500 mt-1">Assignee: {task.assignee}</p>
              </div>
              <div>{getStatusBadge(task.status)}</div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
