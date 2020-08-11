package controllers

import (
	"fmt"
	"myapp/app/models"
	"strconv"

	"github.com/revel/revel"
)

type Movies struct {
	*revel.Controller
}

func (c Movies) List() revel.Result {
	rows, err := models.DB.Query("SELECT * FROM movies")
	if err != nil {
		e := err.Error()
		c.Validation.Error(e).Key("e")
	}
	var res []models.Movie
	for rows.Next() {
		var movie models.Movie
		err = rows.Scan(&movie.ID, &movie.Name, &movie.Comment, &movie.Rating)
		if err != nil {
			fmt.Println(err)
		}
		res = append(res, movie)
	}
	return c.RenderJSON(res)
}

func (c Movies) SearchMovie(movieName string) revel.Result {
	rows, err := models.DB.Query("SELECT * FROM movies WHERE name = $1", movieName)
	if err != nil {
		e := err.Error()
		c.Validation.Error(e).Key("e")
	}
	var res []models.Movie
	for rows.Next() {
		var movie models.Movie
		err = rows.Scan(&movie.ID, &movie.Name, &movie.Comment, &movie.Rating)
		if err != nil {
			fmt.Println(err)
		}
		res = append(res, movie)
	}
	return c.RenderJSON(res)
}

func (c Movies) CreateMovie() revel.Result {
	var res models.Movie

	// Validate Request Data
	name := c.Params.Get("Name")
	c.Validation.Required(name).Key("name").Message("Name of Movie is required")
	rating, err := strconv.Atoi(c.Params.Form.Get("Rating"))
	if err != nil {
		c.Validation.Error("Rating is not of proper datatype")
	}
	comment := c.Params.Form.Get("Comment")
	if c.Validation.HasErrors() {
		fmt.Println("Error")
	}

	res.Name = name
	res.Rating = rating
	res.Comment = comment

	row, err := models.DB.Query("INSERT INTO movies (name, comment, rating) VALUES ($1, $2, $3)",
		res.Name, res.Comment, res.Rating)

	if err != nil {
		e := err.Error()
		c.Validation.Error(e).Key("e")
	}
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.RenderJSON(row)
}

/*func (c App) UserSpecificMovie(movieName string) revel.Result {
	return c.Render(movieName)
}*/

func (c Movies) RateMovie() revel.Result {
	name := c.Params.Get("Name")
	c.Validation.Required(name).Key("name").Message("Name of Movie is required")
	rating := c.Params.Get("Rating")
	c.Validation.Required(rating).Key("rating").Message("Rating for Movie is required")
	movieName := c.Params.Form.Get("Name")
	r, err := strconv.Atoi(rating)
	if err != nil {
		c.Validation.Error("Rating is not of proper datatype")
	}

	row, err := models.DB.Query("UPDATE movies SET rating = $1 WHERE name = $2", r, movieName)
	if err != nil {
		fmt.Println(err)
		e := err.Error()
		c.Validation.Error(e).Key("e")
	}

	return c.RenderJSON(row)
}

func (c Movies) CommentOnMovie() revel.Result {
	movieName := c.Params.Get("Name")
	c.Validation.Required(movieName).Key("movieName").Message("Name of Movie is required")
	comment := c.Params.Get("Comment")
	c.Validation.Required(comment).Key("comment").Message("Comment for Movie is required")

	row, err := models.DB.Query("UPDATE movies SET comment = $1 WHERE name = $2", comment, movieName)
	if err != nil {
		fmt.Println(err)
		e := err.Error()
		c.Validation.Error(e).Key("e")
	}

	return c.RenderJSON(row)
}
