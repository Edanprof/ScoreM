from selenium import webdriver
import os

driver = webdriver.PhantomJS()

driver.get("http://info.nowgoal.com/en/team/summary.aspx?TeamID=59")
driver.execute_script('leftSide(selectTeamID, arrTeam);\
                if (coach.length>0){ mainTitle(teamDetail, coach[0][2 + lang], coach[0][0]);\
                } else { mainTitle(teamDetail,"", 0);\
                }\
                var mainDiv=document.getElementById("div_Table2");\
                if (leagueData.length == 0) {\
                    mainDiv.innerHTML = "no data";\
                } else {\
                    mainDiv.innerHTML = ShowLeaScore() + ShowLeaLetGoal() + ShowBigSmall();\
                }\
                rightSide();')

print driver.execute_script('return document.body.innerHTML;')
