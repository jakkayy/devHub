import * as vscode from 'vscode';

export interface APIContractItem {
    id: string;
    method: string;
    path: string;
    summary: string;
    specVersion: string;
}

export class APIContractTreeItem extends vscode.TreeItem {
    constructor(
        public readonly contract: APIContractItem,
        public readonly collapsibleState: vscode.TreeItemCollapsibleState
    ) {
        super(`${contract.method.toUpperCase()} ${contract.path}`, collapsibleState);
        this.tooltip = `${contract.summary} (${contract.specVersion})`;
        this.description = contract.summary;
        this.iconPath = new vscode.ThemeIcon(
            contract.method.toUpperCase() === 'GET' ? 'symbol-method' : 'symbol-property'
        );
        this.contextValue = 'apiContract';
    }
}

export class APIContractTreeDataProvider implements vscode.TreeDataProvider<APIContractTreeItem> {
    private _onDidChangeTreeData: vscode.EventEmitter<APIContractTreeItem | undefined | null | void> = new vscode.EventEmitter<APIContractTreeItem | undefined | null | void>();
    readonly onDidChangeTreeData: vscode.Event<APIContractTreeItem | undefined | null | void> = this._onDidChangeTreeData.event;

    refresh(): void {
        this._onDidChangeTreeData.fire();
    }

    getTreeItem(element: APIContractTreeItem): vscode.TreeItem {
        return element;
    }

    async getChildren(element?: APIContractTreeItem): Promise<APIContractTreeItem[]> {
        if (element) {
            return [];
        }

        try {
            const res = await fetch('http://localhost:8080/api/v1/contracts');
            const json = (await res.json()) as any;
            const contracts: APIContractItem[] = json.data || [];
            return contracts.map(c => new APIContractTreeItem(c, vscode.TreeItemCollapsibleState.None));
        } catch (err) {
            return [];
        }
    }
}
