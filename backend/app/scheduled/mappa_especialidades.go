package scheduled

import (
	"context"
	"fmt"
)

func get_especialidades(ctx context.Context) error {
	aCtx, err := getActionContext()
	if err != nil {
		return err
	}
	lastLogin := aCtx.Cache.GetLastLogin()
	if lastLogin == nil || !lastLogin.IsValid() {
		return fmt.Errorf("Login not valid")
	}
	aCtx.MappaService.GetEspecialidades()
	return nil
}
