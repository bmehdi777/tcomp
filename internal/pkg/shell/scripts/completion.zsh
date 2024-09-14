_tcomp_show_workspace_files() {
	list_files=""
	for file in ~/.config/tcomp/workspaces/*.yml;
	do
		filename=$(basename $file | cut -d. -f1)
		list_files="${list_files} ${filename}"
	done

	echo $list_files
}

_tcomp_completion() {
	_arguments -C '1:cmd:->cmds' \
		'*:: :->args' \
		&& ret=0

	case "$state" in 
		(cmds)
			local tcomp_commands=(
			'init:print shell script'
			'up:start a tcomp workspace'
			'down:shutdown a tcomp workspace'
			'ls:list all workspaces'
			'new:create a new workspace'
			'version:show version'
			'help:see help'
		)
		_describe -t commands 'command' tcomp_commands && ret=0
		;;
	(args)
		case $line[1] in
			(up|down)
				_arguments "1: :($(_tcomp_show_workspace_files))"
				;;
		esac
		;;
esac

}

compdef '_tcomp_completion' tcomp

