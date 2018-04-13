import time
import os

from selenium import webdriver
from selenium.webdriver.common.desired_capabilities import DesiredCapabilities

time.sleep(15)
host = os.environ.get('HUB_HOST')
game = os.environ.get('GAME_HOST')
print(host, game)
firefox = webdriver.Remote(
          command_executor='http://{}:4444/wd/hub'.format(host),
          desired_capabilities=DesiredCapabilities.FIREFOX) 

print(host)
firefox.get('http://{}/'.format(game))
print(firefox.title)
firefox.quit()