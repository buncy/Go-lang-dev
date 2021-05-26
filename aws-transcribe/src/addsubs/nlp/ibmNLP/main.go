package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/watson-developer-cloud/go-sdk/v2/naturallanguageunderstandingv1"
)

func main() {
	authenticator := &core.IamAuthenticator{
		ApiKey: "LUhNr2bRc3PWQ5tWaik5IzQVc1OxnOdN8TvrCS1fdceT",
	}

	options := &naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
		Version:       core.StringPtr("2020-04-01"),
		Authenticator: authenticator,
	}

	naturalLanguageUnderstanding, naturalLanguageUnderstandingErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(options)

	if naturalLanguageUnderstandingErr != nil {
		panic(naturalLanguageUnderstandingErr)
	}

	naturalLanguageUnderstanding.SetServiceURL("https://api.kr-seo.natural-language-understanding.watson.cloud.ibm.com")

	//url := "www.ibm.com"
	limit := int64(3)
	mentions := true
	explaination := true
	text := "No. Please type a statement for a witness statement addresses. Simple solicitors address one adjust to adjust three and then post good underneath Date is 19 February 2010. Your reference is A. For alpha, B. For bravo, C. For charlie slash P. For papa, Q. For Quebec alpha Romeo slash 12 slash 345 hour references A. For alpha, C. For charlie, C. For charlie. 9876 Inspector of the next line is joe blogs claim technician. On the next line tom smith underneath that, please put a rectangle box with client interviewed in inverted commas. Next line here says we your client Henry jones. Next line, accident data 12 of May 2000 and eight paragraph. We are pleased to confirm that we have interviewed your client and as such, we are pleased to supply our findings to you in the following manner paragraph please type clients underlined and bold as the head in New paragraph. Henry Jones is 32 years of age. He lives with his wife and four Children and is currently unemployed paragraph and gave his evidence quite well. Full stop. He has a confident nature and a polite dispossession for stuff. He answered all questions asked of him and we were not given any particular reason to doubt the evidence he offered throughout the interview process. Full stop. Next heading is bold underline and his evidence. New paragraph. We would refer to the attached statement which you will wish to consider in its entirety. Come up briefly. The main points can be summarized as follows. Next head in a full stop and then accident circumstances. M bold not underlined Next paragraph please bullet point these. He was working for simple signs LTD as a factory operative and had been since the 9th, the 1st 2007. The full stop. The company manufactured road signs and he states he did not receive any formal training before commencing his job full stop. Next bullet point. The incident happened on 12 May 2000 and eight at 10:30 a.m. Is normal job will be to place templates on metal road signs, but this was not required. And so he was told to pull the feet off barriers, barriers are used by construction companies and they needed to be re sprayed yellow full stop. He was told by a supervisor come with john smith come out to pull the barriers off their supporting feet instead of using a sledgehammer come out as it was suggested that a sledgehammer would bring the fees for stop. It was also suggested that a machine could be used, but it was cheaper to use manual labor for stuff. New bullet points. He was on. His third barrier was pulling it out of its feet when he told the shoulder muscles in his left shoulder full stop, no bullet point, he spoke to john smith about his injury, who said he was just trying to get out of the job full stuff. However, when your client stated that he needed to go to hospital, come out, john Clyde, get yourself away, then get yourself away then in inverted thomas please. Mhm New Point, he is not aware if the incident was reported in the accident report work only is not aware of any health and safety executive investigation, full stop. New bullet points. He believed the interesting could have been avoided if he had taken the feet off with a sledgehammer for stop. He also reiterated that the feet would normally be taken off with the machine, but in inverted commas, john smith wanted me to pull them off because it was much cheaper clothes, inverted commas, new bullet points. He states there was no health and safety policy in place at all at all at the time of the incident and he had never signed any health and safety documents will stop. Next head in is be forced up and then Quantum, please put that in bold And then start a new bullet point. He arrived at hospital at 12:30 PM and was seen by a doctor. He was given a support for his shoulder and told to take painkillers full stuff. He was diagnosed as he had torn muscles in his left shoulder and had tissue damage, full stuff. He was told to get a signal from his GP full stuff. Next bullet points ultimately, he got a signature for some three stroke, four weeks for stuff and the bullet point. He still continues to see his G. P. Once each month and is prescribed numerous medications. He has not had any physiotherapy and continues to have problems sleeping and can't do any heavy lifting full stuff. He describes the difficulties he has in cold weather And in connection with his social life and domestic duties and as a result comer, he has been unemployed for the last 13 months. New bullet points, he did get a job as a large over in June 2009, but only lasted four days due to his shoulder injury. Full stop. End of the bullet points. New head in opinion stroke recommendation, which is bold and underline new paragraph, not bullet pointed. The claimant suggests that there was no health and safety policy in force at his place of employment and his and he was never given any manual handling trading. He describes being instructed to pull the barriers out of their supporting feet by hand, but suggests that on occasion this would have been performed by a machine, but it was in inverted commas cheaper to use manual labor full stop. Furthermore, he had been told not to use a sledgehammer to dislodge the feet as it had the potential of damaging them full stuff. New paragraph. If indeed, he had not gone through any manual handling training combat and clearly his employers will have a case to answer. May well struggle to evidence any risk assessment relating to the activities he was undertaking at the material time. Full stop. New heading is time sheet, which is bold underlined under that we have engaged upon this file as follows. Uh please put this in bullet points. Final administration and then tap across uh, then hyphen uh, 40 minutes. Next bullet point Interview with your appliance. Hyphen one hour, 15 minutes. Next bullet point technical compilation, hyphen 40 minutes finalization. Hyphen 20 minutes. Next header is enclosures. Please bolden dilemma. And we enclose the following again bullet points a typed version of the draft statement. We have taken the next bullet point, a note of our feet. Uh Next line if we can be of any further assistance, please do not hesitate to let us know yours faithfully."

	result, detailedResponse, responseErr := naturalLanguageUnderstanding.Analyze(
		&naturallanguageunderstandingv1.AnalyzeOptions{
			Text: &text,
			Features: &naturallanguageunderstandingv1.Features{
				Categories: &naturallanguageunderstandingv1.CategoriesOptions{
					Limit:       &limit,
					Explanation: &explaination,
				},
				Entities: &naturallanguageunderstandingv1.EntitiesOptions{
					Limit:    &limit,
					Mentions: &mentions,
				},
				Keywords: &naturallanguageunderstandingv1.KeywordsOptions{
					Limit: &limit,
				},
				Relations: &naturallanguageunderstandingv1.RelationsOptions{},
				Sentiment: &naturallanguageunderstandingv1.SentimentOptions{},
			},
		},
	)
	if responseErr != nil {
		fmt.Println(detailedResponse)
		panic(responseErr)
	}
	b, _ := json.MarshalIndent(result, "", "   ")
	f, err := os.Create("/home/karthik/go/src/golangdev/aws-transcribe/src/addsubs/nlp/ibmresult.txt")
	if err != nil {
		fmt.Printf("error creating file %v", err)
		return
	}
	f.WriteString(string(b))
	f.Close()
	//fmt.Println(string(b))
}
