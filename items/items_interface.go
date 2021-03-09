package items


import (
	"../animations"
	"../characters"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
)


/*	Vor.: "constant" muss ein gültiger Bezeichner für ein Item sein
	Eff.: Ein Item-Objekt mit dem Bezeichner "constant", 
	das keine Bombe ist, ist geliefert.
	*data erfüllt das Interface Item
	
	NewItem(constant uint8) *data
*/

/*	Vor.: -
	Eff.: Ein Bomben-Objekt ist geliefert.
	*bombe erfüllt das Interface Item
	
	NewBomb() *bombe
*/

type Item interface{
	
	/*	Vor.: -
	 * 	Eff.: Die zum Objekt gehörenden Animation ist geliefert, 
	 * 	in der Animation ist auch die Position des Items 
	 * 	als pixel.Vec gespeichert.
	 */
	Ani() animations.Animation
	
	/*	Vor.: -
	 * 	Eff.: Das Objekt ist entweder zerstörbar, falls true übergeben 
	 * 	wurde oder unzerstörbar, falls false übergeben wurde.
	 */
	SetDestroyable (b bool)
	
	/*	Vor.: -
	 * 	Eff.: true ist genu dann geliefert, wenn das Objekt zerstörbar ist
	 */
	IsDestroyable () bool
	
	/*	Vor.: -
	 * 	Eff.: Das Item ist nun sichtbar genau dann, wenn b=true ist
	 */
	SetVisible(b bool)
	
	/*	Vor.: -
	 * 	Eff.: true ist genu dann geliefert, wenn das Objekt sichtbar ist
	 */
	IsVisible() bool
	
	/*	Vor.: win darf nicht nil sein - das fenser win ist geöffnet
	 * 	Eff.: Das Item ist gezeichnet
	 */
	Draw (win *pixelgl.Window)
	
	/*	Vor.: -
	 * 	Eff.: Das Item wird an die Position nach Vektor pos gesetzt,
	 *  ist aber noch nicht dort gezeichnet: Aufruf von Draw nötig!
	 */
	SetPos (pos pixel.Vec)
	
	/*	Vor.: -
	 * 	Eff.: Der Vektor der Position das Items ist geliefert 
	 */
	GetPos () pixel.Vec
	
}

type Bombe interface {
	Item
	/*	Vor.: -
	 * 	Eff.: Falls das Objekt keinen Besitzer hat ist das Tupel false,nil geliefert.
	 * 	Falls das Objekt einen Besitzer hat, ist true und ein Zeiger auf den Besitzer
	 * 	geliefert.
	 */
	Owner () (bool,characters.Player)
	
	/*	Vor.: Das Objekt hat noch keinen Besitzer.
	 * 	Eff.: Der Besitzer des Objekt ist nun der übergebene player.
	 */
	SetOwner (player characters.Player)
	
	
}
