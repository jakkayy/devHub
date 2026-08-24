'use client';

import React from 'react';
import dynamic from 'next/dynamic';
import TaskWidget from '@/components/TaskWidget';
import GitHubPRWidget from '@/components/GitHubPRWidget';
import APIContractWidget from '@/components/APIContractWidget';
import MiroNodeLinkPanel from '@/components/MiroNodeLinkPanel';
import { Activity, Cpu, Layers, Terminal } from 'lucide-react';

// Dynamically import MiroWidget with SSR disabled for optimal page performance
const MiroWidget = dynamic(() => import('@/components/MiroWidget'), {
  ssr: false,
  loading: () => (
    <div className="bg-zinc-900/90 border border-zinc-800 rounded-xl p-5 shadow-lg flex items-center justify-center min-h-[350px]">
      <span className="text-sm text-zinc-500 animate-pulse">Loading Miro Architecture Board...</span>
    </div>
  ),
});

export default function DashboardPage() {
  return (
    <div className="min-h-screen bg-zinc-950 text-zinc-100 flex flex-col">
      {/* Header */}
      <header className="border-b border-zinc-800 bg-zinc-900/50 backdrop-blur px-6 py-4 flex items-center justify-between sticky top-0 z-50">
        <div className="flex items-center gap-3">
          <div className="p-2 rounded-lg bg-emerald-500/10 border border-emerald-500/20 text-emerald-400">
            <Terminal className="w-5 h-5" />
          </div>
          <div>
            <h1 className="text-lg font-bold tracking-tight text-zinc-100">devHub</h1>
            <p className="text-xs text-zinc-400">Internal Developer Command Center</p>
          </div>
        </div>

        <div className="flex items-center gap-3">
          <span className="flex items-center gap-1.5 text-xs text-emerald-400 bg-emerald-950/60 px-3 py-1 rounded-full border border-emerald-800/40 font-mono">
            <span className="w-2 h-2 rounded-full bg-emerald-400 animate-pulse"></span>
            System Online
          </span>
        </div>
      </header>

      {/* Main Content */}
      <main className="flex-1 max-w-7xl w-full mx-auto p-6 space-y-6">
        {/* Quick Stats Bar */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div className="bg-zinc-900/60 border border-zinc-800/80 p-4 rounded-xl flex items-center gap-4">
            <div className="p-3 rounded-lg bg-blue-500/10 text-blue-400 border border-blue-500/20">
              <Layers className="w-5 h-5" />
            </div>
            <div>
              <p className="text-xs text-zinc-400">Connected Services</p>
              <p className="text-base font-semibold text-zinc-100 mt-0.5">6 Active Tools</p>
            </div>
          </div>

          <div className="bg-zinc-900/60 border border-zinc-800/80 p-4 rounded-xl flex items-center gap-4">
            <div className="p-3 rounded-lg bg-purple-500/10 text-purple-400 border border-purple-500/20">
              <Cpu className="w-5 h-5" />
            </div>
            <div>
              <p className="text-xs text-zinc-400">Backend Engine</p>
              <p className="text-base font-semibold text-zinc-100 mt-0.5">Go (Gin + GORM)</p>
            </div>
          </div>

          <div className="bg-zinc-900/60 border border-zinc-800/80 p-4 rounded-xl flex items-center gap-4">
            <div className="p-3 rounded-lg bg-emerald-500/10 text-emerald-400 border border-emerald-500/20">
              <Activity className="w-5 h-5" />
            </div>
            <div>
              <p className="text-xs text-zinc-400">Caching & DB</p>
              <p className="text-base font-semibold text-zinc-100 mt-0.5">Redis + PostgreSQL</p>
            </div>
          </div>
        </div>

        {/* Miro Architecture & Node-to-Code Mapping Section */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <div className="lg:col-span-2">
            <MiroWidget />
          </div>
          <div>
            <MiroNodeLinkPanel />
          </div>
        </div>

        {/* Dashboard Widgets Grid */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <TaskWidget />
          <GitHubPRWidget />
          <APIContractWidget />
        </div>
      </main>

      {/* Footer */}
      <footer className="border-t border-zinc-800/80 py-4 px-6 text-center text-xs text-zinc-500">
        devHub &copy; {new Date().getFullYear()} - Platform Engineering Workspace
      </footer>
    </div>
  );
}
