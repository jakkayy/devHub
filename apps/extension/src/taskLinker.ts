import * as vscode from 'vscode';

export function registerTaskLinkerCommand(context: vscode.ExtensionContext) {
    const disposable = vscode.commands.registerCommand('devhub.linkSelectionToTask', async () => {
        const editor = vscode.window.activeTextEditor;
        if (!editor) {
            vscode.window.showWarningMessage('No active editor open');
            return;
        }

        const filePath = editor.document.uri.fsPath;
        const line = editor.selection.active.line + 1;

        const taskID = await vscode.window.showInputBox({
            prompt: 'Enter Task ID to link with this code location (e.g. TASK-101)',
            placeHolder: 'TASK-101'
        });

        if (!taskID) {
            return;
        }

        try {
            const res = await fetch('http://localhost:8080/api/v1/links', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    task_id: taskID,
                    file_path: filePath,
                    line_number: line,
                    git_branch: 'develop'
                })
            });

            if (res.ok) {
                vscode.window.showInformationMessage(`Successfully linked line ${line} in ${filePath.split('/').pop()} to ${taskID}!`);
            } else {
                vscode.window.showErrorMessage('Failed to link code location to Task on devHub backend');
            }
        } catch (err: any) {
            vscode.window.showErrorMessage(`Error linking code: ${err.message}`);
        }
    });

    context.subscriptions.push(disposable);
}
