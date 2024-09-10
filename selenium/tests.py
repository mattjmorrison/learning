import unittest
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By


class Tests(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        cls.driver = webdriver.Remote(
           command_executor='http://host.docker.internal:4444/wd/hub',
           options=webdriver.FirefoxOptions()
        )
        cls.driver.set_page_load_timeout(15)

    @classmethod
    def tearDownClass(cls):
        cls.driver.quit()

    def test_home_page_mentions_unix(self):
        self.driver.get('https://www.mattjmorrison.com')
        self.assertIn('Unix', self.driver.page_source)

    def test_home_page_mentions_python(self):
        self.driver.get('https://www.mattjmorrison.com')
        self.assertIn('Python', self.driver.page_source)

    def test_home_page_mentions_javascript(self):
        self.driver.get('https://www.mattjmorrison.com')
        self.assertIn('JavaScript', self.driver.page_source)

    def test_home_page_mentions_vim(self):
        self.driver.get('https://www.mattjmorrison.com')
        self.assertIn('vim', self.driver.page_source)

    def test_home_page_mentions_agile(self):
        self.driver.get('https://www.mattjmorrison.com')
        self.assertIn('agile', self.driver.page_source)

    def test_home_page_mentions_ruby_even_though_it_should_not(self):
        self.driver.get('https://www.mattjmorrison.com')
        self.assertIn('Ruby', self.driver.page_source)


if __name__ == '__main__':
    unittest.main()