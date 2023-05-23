package models

import (
	"fmt"
)

func MakeCommandCd(argument string) Command {
	return Cd{argument}
}

type Cd struct {
	argument string
}

func (Cd) GetCommandType() CommandType {
	return CD
}

func (cd Cd) String() string {
	return fmt.Sprintf("cd %s", cd.argument)
}

func (cd Cd) Copy() Command {
	return MakeCommandCd(cd.argument)
}

func (cd Cd) equal(command Command) bool {
	if other, ok := command.(Cd); ok {
		return cd.argument == other.argument
	}
	return false
}

func (cd Cd) apply(current *Directory) (*Directory, error) {
	switch cd.argument {
	case "/":
		{
			for current != current.parent {
				current = current.parent
			}
			return current, nil
		}
	case "..":
		{
			return current.parent, nil
		}
	default:
		{
			return cd.handleChangeDirectory(current)
		}
	}
}

func (cd Cd) handleChangeDirectory(current *Directory) (*Directory, error) {

	if current.items == nil {
		return nil, directory_content_is_unknown_error(current.name, cd.argument)
	}

	if _, item_exists := current.items[cd.argument]; !item_exists {
		return nil, directory_does_not_exist_error(current.name, cd.argument)
	}

	if directory, ok := current.items[cd.argument].(*Directory); ok {
		return directory, nil
	} else {
		return nil, not_a_directory_error(current.name, cd.argument)
	}
}
