package db

import (
	. "github.com/Serbroda/ragbag/app/pkg/models"
	. "github.com/Serbroda/ragbag/app/pkg/services"
	"github.com/teris-io/shortid"
	"log"
)

const (
	admin        = "admin"
	passwordFile = "adminpassword"
)

func Initialize(us UserService, rs RoleService) {
	res, err := us.FindOneByUsername(admin)
	if exists(res, err) {
		return
	}

	log.Println("Creating admin user")

	pwd := shortid.MustGenerate()
	res, err = us.CreateUser(User{
		Username:  admin,
		Password:  pwd,
		Email:     "admin@admin",
		Active:    true,
		FirstName: "Admin",
		LastName:  "Admin",
	})
	if err != nil {
		panic(err)
	}
	role, err := rs.FindRoleByName("ADMIN")
	if err != nil {
		panic(err)
	}
	err = rs.InsertUserRole(res.ID, role.ID)
}

func exists(user User, err error) bool {
	return user.ID > 0 && err == nil
}

/*func InitializeAdmin(c context.Context, s *services.Services) {
	if s.ExistsUser(c, admin) {
		return
	}

	log.Println("initializing admin user")
	shortId := shortid.MustGenerate()

	user, err := s.CreateUserWithRoles(c, gen.CreateUserParams{
		Username:  admin,
		Password:  shortId,
		Email:     "admin@admin",
		Active:    true,
		FirstName: "Admin",
		LastName:  "Admin",
	}, []string{"ADMIN"})

	if err != nil {
		panic(err.Error())
	}

	_, err = s.CreateSpace(c, gen.CreateSpaceParams{
		ShortID: shortid.MustGenerate(),
		OwnerID: user.ID,
		Name:    "Admin's Space",
	})

	if err != nil {
		panic(err.Error())
	}

	file, err := os.Create(passwordFile)
	if err != nil {
		panic(err.Error())
	}
	//abs, err := filepath.Abs(passwordFile)
	if err == nil {
		//log.Printf("adminpassword file created. Execute 'cat %s' to view initial password.", abs)
		log.Printf("adminpassword file created: %s", shortId)
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s\n", shortId))
	if err != nil {
		panic(err.Error())
	}
}
*/
