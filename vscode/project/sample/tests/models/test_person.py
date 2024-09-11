from django import test
from sample.models.person import Person


class PersonModelTests(test.TestCase):

    def test_can_create_person(self):
        p = Person.objects.create(first_name="Matt", last_name="Morrison")
        self.assertNotEqual(None, p.pk)
        self.assertEqual("Matt", p.first_name)
        self.assertEqual("Morrison", p.last_name)