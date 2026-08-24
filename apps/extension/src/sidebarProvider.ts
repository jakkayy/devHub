import * as vscode from 'vscode';
import { WebviewMessage } from './messaging';

export class DevHubSidebarProvider implements vscode.WebviewViewProvider {
    public static readonly viewType = 'devhub.sidebarView';
    private _view?: vscode.WebviewView;

    constructor(private readonly _extensionUri: vscode.Uri) {}

    public resolveWebviewView(
        webviewView: vscode.WebviewView,
        context: vscode.WebviewViewResolveContext,
        _token: vscode.CancellationToken
    ) {
        this._view = webviewView;

        webviewView.webview.options = {
            enableScripts: true,
            localResourceRoots: [this._extensionUri]
        };

        webviewView.webview.html = this._getHtmlForWebview(webviewView.webview);

        webviewView.webview.onDidReceiveMessage(async (message: WebviewMessage) => {
            switch (message.command) {
                case 'FETCH_TASKS':
                    this._fetchAndSendTasks();
                    break;
                case 'FETCH_CONTRACTS':
                    this._fetchAndSendContracts();
                    break;
                case 'OPEN_DEEP_LINK':
                    if (message.payload?.vscodeUri) {
                        vscode.env.openExternal(vscode.Uri.parse(message.payload.vscodeUri));
                    }
                    break;
                case 'NOTIFY':
                    if (message.payload?.text) {
                        vscode.window.showInformationMessage(message.payload.text);
                    }
                    break;
            }
        });
    }

    private async _fetchAndSendTasks() {
        try {
            const res = await fetch('http://localhost:8080/api/v1/tasks');
            const json = (await res.json()) as any;
            this._view?.webview.postMessage({ type: 'TASKS_LOADED', data: json.data || [] });
        } catch (err: any) {
            this._view?.webview.postMessage({ type: 'ERROR', error: 'Failed to connect to Go Backend' });
        }
    }

    private async _fetchAndSendContracts() {
        try {
            const res = await fetch('http://localhost:8080/api/v1/contracts');
            const json = (await res.json()) as any;
            this._view?.webview.postMessage({ type: 'CONTRACTS_LOADED', data: json.data || [] });
        } catch (err: any) {
            this._view?.webview.postMessage({ type: 'ERROR', error: 'Failed to fetch API specs' });
        }
    }

    private _getHtmlForWebview(webview: vscode.Webview): string {
        return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>devHub Command Center</title>
    <style>
        body {
            font-family: var(--vscode-font-family);
            padding: 12px;
            color: var(--vscode-foreground);
            background-color: var(--vscode-sideBar-background);
        }
        .header {
            font-weight: bold;
            margin-bottom: 12px;
            font-size: 1.1em;
            display: flex;
            align-items: center;
            justify-content: space-between;
            border-bottom: 1px solid var(--vscode-widget-border);
            padding-bottom: 8px;
        }
        .status-badge {
            font-size: 0.75em;
            padding: 2px 6px;
            border-radius: 10px;
            background: rgba(16, 185, 129, 0.2);
            color: #10b981;
            border: 1px solid rgba(16, 185, 129, 0.4);
        }
        .section-title {
            font-size: 0.85em;
            font-weight: 600;
            text-transform: uppercase;
            letter-spacing: 0.5px;
            margin: 14px 0 6px 0;
            color: var(--vscode-descriptionForeground);
        }
        .card-list {
            display: flex;
            flex-direction: column;
            gap: 8px;
        }
        .card {
            background: var(--vscode-editor-background);
            border: 1px solid var(--vscode-widget-border);
            border-radius: 6px;
            padding: 8px 10px;
            font-size: 0.85em;
        }
        .card-header {
            display: flex;
            align-items: center;
            justify-content: space-between;
            font-weight: 600;
            margin-bottom: 4px;
        }
        .tag {
            font-family: var(--vscode-editor-font-family);
            font-size: 0.75em;
            color: #38bdf8;
        }
        .btn {
            background: var(--vscode-button-background);
            color: var(--vscode-button-foreground);
            border: none;
            padding: 5px 10px;
            border-radius: 4px;
            cursor: pointer;
            width: 100%;
            margin-top: 6px;
            font-size: 0.8em;
        }
        .btn:hover {
            background: var(--vscode-button-hoverBackground);
        }
    </style>
</head>
<body>
    <div class="header">
        <span>devHub Command Center</span>
        <span class="status-badge">Connected</span>
    </div>

    <div class="section-title">Active Tasks (Sheets Sync)</div>
    <div id="tasks-container" class="card-list">
        <div class="card">Loading tasks...</div>
    </div>

    <div class="section-title">API Specs (OpenAPI 3.0)</div>
    <div id="contracts-container" class="card-list">
        <div class="card">Loading API contracts...</div>
    </div>

    <script>
        const vscode = acquireVsCodeApi();

        window.addEventListener('message', event => {
            const message = event.data;
            switch (message.type) {
                case 'TASKS_LOADED':
                    renderTasks(message.data);
                    break;
                case 'CONTRACTS_LOADED':
                    renderContracts(message.data);
                    break;
                case 'ERROR':
                    console.error(message.error);
                    break;
            }
        });

        function renderTasks(tasks) {
            const container = document.getElementById('tasks-container');
            if (!tasks || tasks.length === 0) {
                container.innerHTML = '<div class="card">No tasks available</div>';
                return;
            }
            container.innerHTML = tasks.map(t => \`
                <div class="card">
                    <div class="card-header">
                        <span class="tag">\${t.id}</span>
                        <span>\${t.status}</span>
                    </div>
                    <div>\${t.title}</div>
                </div>
            \`).join('');
        }

        function renderContracts(contracts) {
            const container = document.getElementById('contracts-container');
            if (!contracts || contracts.length === 0) {
                container.innerHTML = '<div class="card">No API specs available</div>';
                return;
            }
            container.innerHTML = contracts.map(c => \`
                <div class="card">
                    <div class="card-header">
                        <span class="tag">\${c.method}</span>
                        <span>\${c.path}</span>
                    </div>
                    <div style="font-size: 0.8em; opacity: 0.8;">\${c.summary}</div>
                </div>
            \`).join('');
        }

        // Trigger initial data load
        vscode.postMessage({ command: 'FETCH_TASKS' });
        vscode.postMessage({ command: 'FETCH_CONTRACTS' });
    </script>
</body>
</html>`;
    }
}
