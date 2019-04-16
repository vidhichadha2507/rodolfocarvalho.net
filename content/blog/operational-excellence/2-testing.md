---
title: "Automated Tests as the Basis for Operational Excellence"
slug: "operational-excellence/automated-tests-as-the-basis-for-operational-excellence"
description: >-
  It takes multiple layers of automated testing to operate a successful service
  sustainably. Good automated tests are a fundamental step to achieve
  Operational Excellence and a prerequisite to be able to follow other
  practices. This post discusses multiple approaches to testing and how they
  support the development of new features and the operation of the system.
categories: []
tags: []
series: ["operational-excellence"]
date: 2019-04-22T11:00:00+01:00
---

One of the key ideas behind Operational Excellence is automation.

Automated tests are the basis for Operational Excellence because all of the
other practices assume we have a sufficient set of automated tests to back us up
against regressions and to detect problems as early as possible.

Continuous Integration makes no sense without tests and would be very laborious
and error prone without automated tests.

Continuous Deployment/Delivery would be extremely risky without automated tests.
No one wants to change production software without a sense of how safe it is to
do so.

Software projects benefit from test automation not only because computers are
better than us at repeating tasks. More importantly, by writing down test cases
as code we can unambiguously **communicate intent with other humans**, more
easily reproduce test results and track changes over time.

Here we'll consider four complementary types of tests:

- [Unit]({{< relref "#unit-testing" >}})
- [Integration]({{< relref "#integration-testing" >}})
- [End-to-end]({{< relref "#end-to-end-testing" >}})
- [Canaries]({{< relref "#testing-canaries" >}})

Since automated tests are nothing new, much has been written about them. There
is a lot of material online and offline about unit, integration and end-to-end
tests.

What I consider to be a less known and less commonly implemented practice is
that of [testing canaries]({{< relref "#testing-canaries" >}}). It is, however,
a type of testing that can provide a lot of value to cloud-native software.
Therefore, I'll devote more words to the latter.

## Unit Testing

According to a [2019 survey from
StackOverflow](https://insights.stackoverflow.com/survey/2019) with nearly
90,000 developers:

> [...] they are overwhelmingly in favor of unit testing, whether they currently
> use them or not. In fact, developers at companies who embrace unit testing
> also have slightly higher job satisfaction.

Unit tests are at the bottom of the test pyramid.[^pyramid1][^pyramid2]

[^pyramid1]: The pyramid metaphor is documented in [Test Pyramid by Martin Fowler](https://martinfowler.com/bliki/TestPyramid.html)
[^pyramid2]: A more in-depth discussion is available at [The Practical Test Pyramid by Ham Vocke](https://martinfowler.com/articles/practical-test-pyramid.html)

They are there to document how pieces of your code base should work, to catch
low level problems early during development, and to prevent the regression of
previously fixed bugs.

Your team will typically write lots of unit tests, and ship them with new
functionality and bug fixes.

## Integration Testing

Integration tests are higher level tests that typically operate on more than one
component of the system. For example, for an online store, imagine a product
catalog microservice being tested with only a subset of the components that make
up the store, with a focus on the contract of that microservice.

They go in the middle of the testing pyramid.

## End-to-end Testing

End-to-end tests focus on usage flows of the system as a whole. In the online
store example, end-to-end tests would check that a new user can create an
account and that existing users can find and shop products.

They are often harder to write and slower to run. You want to have fewer of
them.

## Testing Canaries

Not to be confused with *canary deployments*,[^canary-deploy] testing canaries
are programs that continuously live test the most important operations that
users can perform with your software. They are bots that constantly throw
traffic at your system to perform and verify the outcome of key user actions.

[^canary-deploy]: The idea is explained in the article [Canary Release by Danilo Sato](https://martinfowler.com/bliki/CanaryRelease.html), and is in fact complementary to having the test canaries as described above.

![A yellow bird](https://source.unsplash.com/o0S-0Pa4F2M/320x160)

They should be the first ones to detect inconsistent behavior due to an
incorrect deployment, software bug, bad configuration or other problems that
reveal themselves (sometimes only) in the production environment.

Canaries are independent auxiliary components and thus may, and should, be
deployed by their own dedicated deployment pipeline, independent of the pipeline
for the software under test.

Good canaries are simple and easy to understand. Each canary should focus on one
key use case of your software. Have multiple canaries to cover your most
important functionality.

One modern way to implement canaries is to write small programs that run on a
schedule with AWS Lambda.

Don't forget that canaries cost you money. Make the best use of them. They
should be at the tip of your testing pyramid. You shall only have a reasonably
small number of canaries, covering the most critical functionality.

Make sure canaries produce meaningful logs as they can be of help estimating
failure impact, for example "API `/xyz` was down between 11:35 and 11:45, as
demonstrated by our canary logs".

You may want to filter traffic from a canary. Make your canaries identify
themselves by, for example, sending an appropriate `User-Agent` HTTP header.

Integrate your canaries with your monitoring and alerting systems. Whenever a
canary runs into a problem, make sure your alerting system pages in a human
operator to investigate. Be careful with over alerting, though. Remember we are
working with a distributed system and design your canary to tolerate small
problems: delays, missed events, etc. You do not want to wake up a human
operator in the middle of the night for no good reason.

---

One quick note before you go write your tests and fancy canaries.

Umer Mansoor wrote smart [cautionary words about trying to do things you don't
need just because you saw them
elsewhere](https://codeahoy.com/2017/08/19/yagni-cargo-cult-and-overengineering-the-planes-wont-land-just-because-you-built-a-runway-in-your-backyard/).

I'll emphasize here that not all teams need to implement all practices to
achieve Operational Excellence. In fact, each company and team, as they mature,
should concentrate on developing their people and their practices with a focus
on what is necessary and what works for them at each stage of their growth.
