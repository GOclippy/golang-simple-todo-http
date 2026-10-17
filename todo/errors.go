package todo

import "errors"

var ErrTaskAleradyExists = errors.New("ErrTaskAlreadyExist")
var ErrTaskNotFound = errors.New("ErrTaskNotFound")
var ErrNameIsEmpty = errors.New("ErrNameIsEmpty")
var ErrDescriptionIsEmpty = errors.New("ErrDescriptionIsEmpty")
var ErrNameAndDescriptionIsEmpty = errors.New("ErrNameAndDescriptionIsEmpty")