<!DOCTYPE html>
<html lang="es">
  <head>
    <title>AppdeGato</title>
    <style> 
    h2 {
      color: blue;}</style>
      h3 {color: black;}
  </head>
    <body>
      
        <h2>Appdegatos</h2>
        
         <main>
           <h3>imágenes de gatos</h3>
           <!--TODO: para agregar enlace a imagenes de-->
           <p> <a href="#" target="_blank" rel="noopener noreferrer">haz click aqui para ver imagenes de gatos</a></p>

           


           <a href="https://www.freecodecamp.org/espanol" target="_blank" rel="noopener noreferrer"><img src="https://bit.ly/fcc-relaxing-cat" alt="un lindo gato naranja recostad osobre su espalda"></a>
            
           <hr>
          

           <h3>Listas de gatos</h3>
           <div>
           <p>cosas que los gatos <em>aman</em>:
           </p>
          <ul>
            <li>menta gatuna</li>
            <li>apuntadores laser</li>
            <li>lasaña</li>
        
          </ul>
        </div>
          <img src="imagenes/descarga.jpg" alt="lasaña">  
          <p>cosas que los gatos <strong>odian</strong>:</p>
          <ol>
            <li>tratamientos antipulgas</li>
            <li>truenos</li>
            <li>gatos</li>
          </ol>
          <img src="imagenes/gatitos.jpg">
<!--boton de radio-->
          <form action="/enviar-respuesta">
            <label for="interior"><input value="interior" id="interior" type="radio" name="interior-exterior" checked> Interiores</label>

            <br>

            <label for="exterior"><value="exterior" input id="exterior" type="radio" name="interior-exterior"> Exteriores</label>
            <br>
          <!--boton de inerio exterior-->

            <input type="text" placeholder= "foto de tu gato" required >
            <br>
            <label for="carinoso">
              <input id="carinoso" type="checkbox" name="personalidad" checked > cariñoso</label>

            <label for="perezoso"> 
              <input id="perezoso" type="checkbox" name="personalidad"> perezoso</label>
            <label for="energico"> 
              <input id="energico" type="checkbox" name="personalidad"> energico</label>
<br>
            <button type="submit"> enviar </button>
<!--casillas deverificacion-->
<br>
           
            
          </form>

         </main>
         <p><small><footer>sin derechos de autor <a href="https://www.google.com" target="_blank" rel="nonopener noreferrer"> asdasdasdsa</footer></a><small><p>
      
    </body>
</html>
