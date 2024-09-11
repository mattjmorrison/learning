import json
from django import test
from sample.models.person import Person

class PersonAPITests(test.TestCase):

    def test_gets_list_of_all_people(self):
        Person.objects.create(first_name='Matt', last_name='Morrison')
        response = self.client.get('/sample/people/')
        self.assertEqual(200, response.status_code)
        json_response = json.loads(response.content)
        self.assertEqual(1, len(json_response))

