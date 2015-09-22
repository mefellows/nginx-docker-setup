package au.com.onegeek.nginxperf

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

  class Nginx extends Simulation {

    val httpConf = http
      .baseURL("http://api.foo.com")
      .acceptHeader("text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
      .doNotTrackHeader("1")
      .acceptLanguageHeader("en-US,en;q=0.5")
      .acceptEncodingHeader("gzip, deflate")
      .userAgentHeader("Mozilla/5.0 (Windows NT 5.1; rv:31.0) Gecko/20100101 Firefox/31.0")

    val scn = scenario("BasicSimulation")
      .repeat(10) {
        exec(http("request_1")
        .get("/header/host"))
      }

    setUp(
      scn.inject(rampUsers(10000) over (1 seconds))
    ).protocols(httpConf)
  }