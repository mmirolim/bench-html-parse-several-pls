import java.io.PrintWriter
import scala.collection.JavaConversions._
import org.jsoup.Jsoup
import org.json4s.jackson.JsonMethods._
import org.json4s.JsonDSL._
import scala.util.{Failure, Success, Try}

object Parser {

  var fname = "parse-pcuz"

  def extract(url: String) = Try(Jsoup.connect(url).get()) match {
    case Success(d) =>

      val start = System.nanoTime()

      // extract names
      val as = d.getElementsByAttributeValue("style", "font-size:11pt; text-decoration:none;")
      val names = as.map(el => el.text())

      // extract tels
      val abouts = d.getElementsByAttributeValue("class", "line_about")
      val tels = abouts.map(el => el.getElementsByTag("span").get(1).text())

      // zip names to tels and put to maps
      val orgs = (names zip tels) map (e => Map("name" -> e._1, "tel" -> e._2))

      // json encode
      val json = compact(render(orgs))

      // save to file
      val writer = new PrintWriter(fname + "-scala.json")
      writer.write(json)
      writer.close()

      // return time taken in ms
      ((System.nanoTime() - start) / 10E6) + " ms"

    case Failure(e) => println(e)
  }

}
// get whole list in one go
val url = "http://www.pc.uz/trade/orgs/cat1013?sort=0&limit=10000"

println(Parser extract url)

